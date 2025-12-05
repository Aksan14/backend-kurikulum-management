package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CPLController struct {
	cplService service.CPLService
}

func NewCPLController(cplService service.CPLService) *CPLController {
	return &CPLController{cplService: cplService}
}

// GetAllCPL godoc
// @Summary Get all CPL
// @Description Get all CPL dengan pagination
// @Tags CPL
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by kode or nama"
// @Param status query string false "Filter by status"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /cpl [get]
func (c *CPLController) GetAllCPL(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := dto.CPLListRequest{
		Page:      page,
		Limit:     limit,
		Search:    ctx.Query("search"),
		Status:    ctx.Query("status"),
		SortBy:    ctx.Query("sort_by"),
		SortOrder: ctx.Query("sort_order"),
	}

	response, err := c.cplService.GetAllCPL(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPL",
		Data:    response,
	})
}

// GetCPLByID godoc
// @Summary Get CPL by ID
// @Description Get detail CPL berdasarkan ID
// @Tags CPL
// @Produce json
// @Security BearerAuth
// @Param id path string true "CPL ID"
// @Success 200 {object} dto.APIResponse{data=dto.CPLResponse}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cpl/{id} [get]
func (c *CPLController) GetCPLByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.cplService.GetCPLByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPL",
		Data:    response,
	})
}

// CreateCPL godoc
// @Summary Create new CPL
// @Description Create new CPL (Kaprodi only)
// @Tags CPL
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateCPLRequest true "Create CPL Request"
// @Success 201 {object} dto.APIResponse{data=dto.CPLResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /cpl [post]
func (c *CPLController) CreateCPL(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var req dto.CreateCPLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.cplService.CreateCPL(userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "CPL berhasil dibuat",
		Data:    response,
	})
}

// UpdateCPL godoc
// @Summary Update CPL
// @Description Update data CPL (Kaprodi only)
// @Tags CPL
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "CPL ID"
// @Param request body dto.UpdateCPLRequest true "Update CPL Request"
// @Success 200 {object} dto.APIResponse{data=dto.CPLResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /cpl/{id} [put]
func (c *CPLController) UpdateCPL(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UpdateCPLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.cplService.UpdateCPL(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "CPL berhasil diupdate",
		Data:    response,
	})
}

// DeleteCPL godoc
// @Summary Delete CPL
// @Description Delete CPL (Kaprodi only)
// @Tags CPL
// @Produce json
// @Security BearerAuth
// @Param id path string true "CPL ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /cpl/{id} [delete]
func (c *CPLController) DeleteCPL(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.cplService.DeleteCPL(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "CPL berhasil dihapus",
	})
}

// UpdateCPLStatus godoc
// @Summary Update CPL status
// @Description Update status CPL (Kaprodi only)
// @Tags CPL
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "CPL ID"
// @Param request body dto.UpdateCPLStatusRequest true "Update Status Request"
// @Success 200 {object} dto.APIResponse{data=dto.CPLResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /cpl/{id}/status [patch]
func (c *CPLController) UpdateCPLStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UpdateCPLStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.cplService.UpdateCPLStatus(id, req.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Status CPL berhasil diupdate",
		Data:    response,
	})
}

// GetCPLStatistics godoc
// @Summary Get CPL statistics
// @Description Get CPL statistics count by status
// @Tags CPL
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /cpl/statistics [get]
func (c *CPLController) GetCPLStatistics(ctx *gin.Context) {
	response, err := c.cplService.GetCPLStatistics()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan statistik CPL",
		Data:    response,
	})
}

// GetActiveCPL godoc
// @Summary Get active CPL
// @Description Get all CPL with active status
// @Tags CPL
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse{data=[]dto.CPLResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /cpl/active [get]
func (c *CPLController) GetActiveCPL(ctx *gin.Context) {
	response, err := c.cplService.GetActiveCPL()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPL aktif",
		Data:    response,
	})
}
