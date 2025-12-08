package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CPLMKMappingController struct {
	mappingService service.CPLMKMappingService
}

func NewCPLMKMappingController(mappingService service.CPLMKMappingService) *CPLMKMappingController {
	return &CPLMKMappingController{mappingService: mappingService}
}

// GetAllMappings godoc
// @Summary Get all CPL-MK mappings
// @Description Get all CPL-MK mappings with pagination and filters
// @Tags CPL-MK Mapping
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param cpl_id query string false "Filter by CPL ID"
// @Param mata_kuliah_id query string false "Filter by Mata Kuliah ID"
// @Param level query string false "Filter by level (tinggi, sedang, rendah)"
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse
// @Router /cpl-mk-mappings [get]
func (c *CPLMKMappingController) GetAllMappings(ctx *gin.Context) {
	var req dto.CPLMKMappingListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Message: "Parameter tidak valid",
			Error:   err.Error(),
		})
		return
	}

	result, err := c.mappingService.GetAll(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Message: "Gagal mengambil data mapping",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data mapping berhasil diambil",
		Data:    result,
	})
}

// GetMappingByID godoc
// @Summary Get mapping by ID
// @Description Get a specific CPL-MK mapping by ID
// @Tags CPL-MK Mapping
// @Accept json
// @Produce json
// @Param id path string true "Mapping ID"
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse
// @Router /cpl-mk-mappings/{id} [get]
func (c *CPLMKMappingController) GetMappingByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.mappingService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data mapping berhasil diambil",
		Data:    result,
	})
}

// UpsertMapping godoc
// @Summary Create or update CPL-MK mapping
// @Description Create a new mapping or update existing one based on cpl_id and mata_kuliah_id
// @Tags CPL-MK Mapping
// @Accept json
// @Produce json
// @Param request body dto.UpsertCPLMKMappingRequest true "Upsert Mapping Request"
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse "Mapping updated"
// @Success 201 {object} dto.APIResponse "Mapping created"
// @Router /cpl-mk-mappings/upsert [post]
func (c *CPLMKMappingController) UpsertMapping(ctx *gin.Context) {
	var req dto.UpsertCPLMKMappingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	result, isNew, err := c.mappingService.Upsert(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "CPL tidak ditemukan" || err.Error() == "Mata Kuliah tidak ditemukan" {
			statusCode = http.StatusNotFound
		}
		ctx.JSON(statusCode, dto.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	status := http.StatusOK
	message := "Mapping berhasil diupdate"
	if isNew {
		status = http.StatusCreated
		message = "Mapping berhasil dibuat"
	}

	ctx.JSON(status, dto.APIResponse{
		Success: true,
		Message: message,
		Data:    result,
	})
}

// DeleteMapping godoc
// @Summary Delete CPL-MK mapping
// @Description Delete a CPL-MK mapping by ID
// @Tags CPL-MK Mapping
// @Accept json
// @Produce json
// @Param id path string true "Mapping ID"
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse
// @Router /cpl-mk-mappings/{id} [delete]
func (c *CPLMKMappingController) DeleteMapping(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.mappingService.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, dto.APIResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Mapping berhasil dihapus",
	})
}

func (c *CPLMKMappingController) GetCPLsByMataKuliahID(ctx *gin.Context) {
	mataKuliahID := ctx.Query("mata_kuliah_id")
	if mataKuliahID == "" {
		ctx.JSON(http.StatusBadRequest, dto.APIResponse{
			Success: false,
			Message: "mata_kuliah_id parameter is required",
		})
		return
	}

	result, err := c.mappingService.GetCPLsByMataKuliahID(mataKuliahID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.APIResponse{
			Success: false,
			Message: "Gagal mengambil data CPL untuk mata kuliah",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data CPL berhasil diambil",
		Data:    result,
	})
}