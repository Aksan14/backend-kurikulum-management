package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MataKuliahController struct {
	mkService service.MataKuliahService
}

func NewMataKuliahController(mkService service.MataKuliahService) *MataKuliahController {
	return &MataKuliahController{mkService: mkService}
}

// @Router /mata-kuliah [get]
func (c *MataKuliahController) GetAllMataKuliah(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	semester, _ := strconv.Atoi(ctx.Query("semester"))

	req := dto.MataKuliahListRequest{
		Page:      page,
		Limit:     limit,
		Search:    ctx.Query("search"),
		Semester:  semester,
		Jenis:     ctx.Query("jenis"),
		Status:    ctx.Query("status"),
		SortBy:    ctx.Query("sort_by"),
		SortOrder: ctx.Query("sort_order"),
	}

	response, err := c.mkService.GetAllMataKuliah(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data mata kuliah berhasil diambil",
		Data:    response,
	})
}

func (c *MataKuliahController) GetMataKuliahByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.mkService.GetMataKuliahByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data mata kuliah berhasil diambil",
		Data:    response,
	})
}

// GetMyMataKuliah godoc
// @Summary Get my Mata Kuliah
// @Description Get all Mata Kuliah assigned to current user (Dosen)
// @Tags MataKuliah
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /mata-kuliah/my [get]
func (c *MataKuliahController) GetMyMataKuliah(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	response, err := c.mkService.GetMyMataKuliah(userID.(string), page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data mata kuliah berhasil diambil",
		Data:    response,
	})
}

// GetMataKuliahBySemester godoc
// @Summary Get Mata Kuliah by semester
// @Description Get all Mata Kuliah for specific semester
// @Tags MataKuliah
// @Produce json
// @Security BearerAuth
// @Param semester path int true "Semester number (1-8)"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /mata-kuliah/semester/{semester} [get]
func (c *MataKuliahController) GetMataKuliahBySemester(ctx *gin.Context) {
	semester, err := strconv.Atoi(ctx.Param("semester"))
	if err != nil || semester < 1 || semester > 8 {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Semester tidak valid (harus 1-8)",
		})
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	response, err := c.mkService.GetMataKuliahBySemester(semester, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data mata kuliah semester " + strconv.Itoa(semester) + " berhasil diambil",
		Data:    response,
	})
}

// GetMataKuliahByDosen godoc
// @Summary Get Mata Kuliah by dosen
// @Description Get all Mata Kuliah assigned to specific dosen
// @Tags MataKuliah
// @Produce json
// @Security BearerAuth
// @Param dosen_id path string true "Dosen ID"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /mata-kuliah/dosen/{dosen_id} [get]
func (c *MataKuliahController) GetMataKuliahByDosen(ctx *gin.Context) {
	dosenID := ctx.Param("dosen_id")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	response, err := c.mkService.GetMataKuliahByDosen(dosenID, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Data mata kuliah dosen berhasil diambil",
		Data:    response,
	})
}

// CreateMataKuliah godoc
// @Summary Create new Mata Kuliah
// @Description Create new Mata Kuliah (Kaprodi only)
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateMataKuliahRequest true "Create Mata Kuliah Request"
// @Success 201 {object} dto.APIResponse{data=dto.MataKuliahResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Router /mata-kuliah [post]
func (c *MataKuliahController) CreateMataKuliah(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var req dto.CreateMataKuliahRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.mkService.CreateMataKuliah(userID.(string), req)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "kode mata kuliah sudah digunakan" {
			status = http.StatusConflict
		}
		ctx.JSON(status, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Mata kuliah berhasil dibuat",
		Data:    response,
	})
}

// UpdateMataKuliah godoc
// @Summary Update Mata Kuliah
// @Description Update data Mata Kuliah (Kaprodi only)
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Mata Kuliah ID"
// @Param request body dto.UpdateMataKuliahRequest true "Update Mata Kuliah Request"
// @Success 200 {object} dto.APIResponse{data=dto.MataKuliahResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /mata-kuliah/{id} [put]
func (c *MataKuliahController) UpdateMataKuliah(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UpdateMataKuliahRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.mkService.UpdateMataKuliah(id, req)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "mata kuliah tidak ditemukan" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Mata kuliah berhasil diupdate",
		Data:    response,
	})
}

// DeleteMataKuliah godoc
// @Summary Delete Mata Kuliah
// @Description Soft delete Mata Kuliah (Kaprodi only)
// @Tags MataKuliah
// @Produce json
// @Security BearerAuth
// @Param id path string true "Mata Kuliah ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /mata-kuliah/{id} [delete]
func (c *MataKuliahController) DeleteMataKuliah(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.mkService.DeleteMataKuliah(id); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "mata kuliah tidak ditemukan" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Mata kuliah berhasil dihapus",
	})
}

// ToggleMataKuliahStatus godoc
// @Summary Toggle Mata Kuliah status
// @Description Toggle Mata Kuliah between aktif/nonaktif
// @Tags MataKuliah
// @Produce json
// @Security BearerAuth
// @Param id path string true "Mata Kuliah ID"
// @Success 200 {object} dto.APIResponse{data=dto.MataKuliahResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /mata-kuliah/{id}/toggle-status [patch]
func (c *MataKuliahController) ToggleMataKuliahStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.mkService.ToggleMataKuliahStatus(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Status mata kuliah berhasil diubah",
		Data:    response,
	})
}

// AssignDosen godoc
// @Summary Assign dosen to Mata Kuliah
// @Description Assign dosen pengampu or koordinator to Mata Kuliah (Kaprodi only)
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Mata Kuliah ID"
// @Param request body dto.AssignDosenRequest true "Assign Dosen Request"
// @Success 200 {object} dto.APIResponse{data=dto.MataKuliahResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /mata-kuliah/{id}/assign-dosen [patch]
func (c *MataKuliahController) AssignDosen(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.AssignDosenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	// Validate at least one field is provided
	if req.DosenPengampuID == nil && req.KoordinatorID == nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Minimal satu field harus diisi: dosen_pengampu_id atau koordinator_id",
		})
		return
	}

	response, err := c.mkService.AssignDosen(id, req)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "mata kuliah tidak ditemukan" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Dosen berhasil di-assign ke mata kuliah",
		Data:    response,
	})
}

// UnassignDosen godoc
// @Summary Unassign dosen from Mata Kuliah
// @Description Remove dosen pengampu or koordinator from Mata Kuliah (Kaprodi only)
// @Tags MataKuliah
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Mata Kuliah ID"
// @Param request body dto.UnassignDosenRequest true "Unassign Dosen Request"
// @Success 200 {object} dto.APIResponse{data=dto.MataKuliahResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /mata-kuliah/{id}/unassign-dosen [patch]
func (c *MataKuliahController) UnassignDosen(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UnassignDosenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	// Default to "all" if type is empty
	dosenType := req.Type
	if dosenType == "" {
		dosenType = "all"
	}

	// Validate type
	if dosenType != "pengampu" && dosenType != "koordinator" && dosenType != "all" {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Type harus salah satu dari: pengampu, koordinator, all",
		})
		return
	}

	response, err := c.mkService.UnassignDosen(id, dosenType)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "mata kuliah tidak ditemukan" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Dosen berhasil di-unassign dari mata kuliah",
		Data:    response,
	})
}
