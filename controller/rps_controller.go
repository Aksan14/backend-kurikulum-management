package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RPSController struct {
	rpsService service.RPSService
}

func NewRPSController(rpsService service.RPSService) *RPSController {
	return &RPSController{rpsService: rpsService}
}

// GetAllRPS godoc
// @Summary Get all RPS
// @Description Get all RPS dengan pagination
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param mata_kuliah_id query string false "Filter by mata kuliah"
// @Param dosen_id query string false "Filter by dosen"
// @Param status query string false "Filter by status"
// @Param tahun_ajaran query string false "Filter by tahun ajaran"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /rps [get]
func (c *RPSController) GetAllRPS(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := dto.RPSListRequest{
		Page:          page,
		Limit:         limit,
		MataKuliahID:  ctx.Query("mata_kuliah_id"),
		DosenID:       ctx.Query("dosen_id"),
		Status:        ctx.Query("status"),
		TahunAkademik: ctx.Query("tahun_ajaran"),
		SortBy:        ctx.Query("sort_by"),
		SortOrder:     ctx.Query("sort_order"),
	}

	response, err := c.rpsService.GetAllRPS(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data RPS",
		Data:    response,
	})
}

// GetRPSByID godoc
// @Summary Get RPS by ID
// @Description Get detail RPS berdasarkan ID
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=dto.RPSResponse}
// @Failure 404 {object} dto.ErrorResponse
// @Router /rps/{id} [get]
func (c *RPSController) GetRPSByID(ctx *gin.Context) {
	id := ctx.Param("rps_id")

	response, err := c.rpsService.GetRPSByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data RPS",
		Data:    response,
	})
}

// CreateRPS godoc
// @Summary Create new RPS
// @Description Create new RPS (Dosen)
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.RPSRequest true "Create RPS Request"
// @Success 201 {object} dto.APIResponse{data=dto.RPSResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps [post]
func (c *RPSController) CreateRPS(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	userName, _ := ctx.Get("userName")

	var req dto.RPSRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	dosenNama := ""
	if userName != nil {
		dosenNama = userName.(string)
	}

	response, err := c.rpsService.CreateRPS(userID.(string), dosenNama, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "RPS berhasil dibuat",
		Data:    response,
	})
}

// UpdateRPS godoc
// @Summary Update RPS
// @Description Update data RPS (Dosen)
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "RPS ID"
// @Param request body dto.UpdateRPSRequest true "Update RPS Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{id} [put]
func (c *RPSController) UpdateRPS(ctx *gin.Context) {
	id := ctx.Param("rps_id")

	var req dto.UpdateRPSRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsService.UpdateRPSPartial(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "RPS berhasil diupdate",
		Data:    response,
	})
}

// DeleteRPS godoc
// @Summary Delete RPS
// @Description Delete RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{id} [delete]
func (c *RPSController) DeleteRPS(ctx *gin.Context) {
	id := ctx.Param("rps_id")

	if err := c.rpsService.DeleteRPS(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "RPS berhasil dihapus",
	})
}

// SubmitRPS godoc
// @Summary Submit RPS for approval
// @Description Submit RPS untuk di-review Kaprodi
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=dto.RPSResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{id}/submit [patch]
func (c *RPSController) SubmitRPS(ctx *gin.Context) {
	id := ctx.Param("rps_id")

	response, err := c.rpsService.SubmitRPS(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "RPS berhasil disubmit untuk review",
		Data:    response,
	})
}

// ApproveRPS godoc
// @Summary Approve RPS
// @Description Approve RPS (Kaprodi only)
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "RPS ID"
// @Param request body dto.ApproveRPSRequest true "Approve RPS Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{id}/approve [patch]
func (c *RPSController) ApproveRPS(ctx *gin.Context) {
	id := ctx.Param("rps_id")
	userID, _ := ctx.Get("userID")

	var req dto.ApproveRPSRequest
	ctx.ShouldBindJSON(&req)

	var catatan *string
	if req.Catatan != "" {
		catatan = &req.Catatan
	}

	response, err := c.rpsService.ApproveRPS(id, userID.(string), catatan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "RPS berhasil diapprove",
		Data:    response,
	})
}

// RejectRPS godoc
// @Summary Reject RPS
// @Description Reject RPS (Kaprodi only)
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "RPS ID"
// @Param request body dto.RejectRPSRequest true "Reject RPS Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{id}/reject [patch]
func (c *RPSController) RejectRPS(ctx *gin.Context) {
	id := ctx.Param("rps_id")
	userID, _ := ctx.Get("userID")

	var req dto.RejectRPSRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Alasan penolakan wajib diisi",
			Error:   err.Error(),
		})
		return
	}

	alasan := &req.Alasan
	response, err := c.rpsService.RejectRPS(id, userID.(string), alasan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "RPS berhasil ditolak",
		Data:    response,
	})
}

// RequestRevision godoc
// @Summary Request revision for RPS
// @Description Request revision for RPS (Kaprodi only)
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "RPS ID"
// @Param request body dto.RequestRevisionRequest true "Request Revision"
// @Success 200 {object} dto.APIResponse{data=dto.RPSResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{id}/request-revision [patch]
func (c *RPSController) RequestRevision(ctx *gin.Context) {
	id := ctx.Param("rps_id")
	userID, _ := ctx.Get("userID")

	var req dto.RequestRevisionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Catatan revisi wajib diisi",
			Error:   err.Error(),
		})
		return
	}

	catatan := &req.Catatan
	response, err := c.rpsService.RequestRevision(id, userID.(string), catatan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "RPS membutuhkan revisi",
		Data:    response,
	})
}

// GetMyRPS godoc
// @Summary Get my RPS
// @Description Get all RPS for current user (Dosen)
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param status query string false "Filter by status"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /rps/my [get]
func (c *RPSController) GetMyRPS(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	status := ctx.Query("status")

	response, err := c.rpsService.GetRPSByDosen(userID.(string), status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data RPS",
		Data:    response,
	})
}

// GetRPSByMataKuliah godoc
// @Summary Get RPS by Mata Kuliah
// @Description Get all RPS for specific Mata Kuliah
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param mata_kuliah_id path string true "Mata Kuliah ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /rps/mata-kuliah/{mata_kuliah_id} [get]
func (c *RPSController) GetRPSByMataKuliah(ctx *gin.Context) {
	mkID := ctx.Param("mata_kuliah_id")

	response, err := c.rpsService.GetRPSByMataKuliah(mkID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data RPS",
		Data:    response,
	})
}

// ========== CPMK Endpoints ==========

// AddCPMK godoc
// @Summary Add CPMK to RPS
// @Description Add CPMK (Capaian Pembelajaran Mata Kuliah) to RPS
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Param request body dto.RPSCPMKRequest true "Create CPMK Request"
// @Success 201 {object} dto.APIResponse{data=dto.RPSCPMKResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/cpmk [post]
func (c *RPSController) AddCPMK(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	var req dto.RPSCPMKRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsService.AddCPMK(rpsID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "CPMK berhasil ditambahkan",
		Data:    response,
	})
}

// UpdateCPMK godoc
// @Summary Update CPMK
// @Description Update CPMK data
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cpmk_id path string true "CPMK ID"
// @Param request body dto.RPSCPMKRequest true "Update CPMK Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSCPMKResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/cpmk/{cpmk_id} [put]
func (c *RPSController) UpdateCPMK(ctx *gin.Context) {
	cpmkID := ctx.Param("cpmk_id")

	var req dto.RPSCPMKRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsService.UpdateCPMK(cpmkID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "CPMK berhasil diupdate",
		Data:    response,
	})
}

// DeleteCPMK godoc
// @Summary Delete CPMK
// @Description Delete CPMK from RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param cpmk_id path string true "CPMK ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/cpmk/{cpmk_id} [delete]
func (c *RPSController) DeleteCPMK(ctx *gin.Context) {
	cpmkID := ctx.Param("cpmk_id")

	if err := c.rpsService.DeleteCPMK(cpmkID); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "CPMK berhasil dihapus",
	})
}

// GetAllCPMK godoc
// @Summary Get All CPMK
// @Description Get all CPMK from all RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSCPMKResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/cpmk [get]
func (c *RPSController) GetAllCPMK(ctx *gin.Context) {
	response, err := c.rpsService.GetAllCPMK()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan semua data CPMK",
		Data:    response,
	})
}

// GetCPMKByRPS godoc
// @Summary Get CPMK by RPS ID
// @Description Get all CPMK for a specific RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSCPMKResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/cpmk [get]
func (c *RPSController) GetCPMKByRPS(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	response, err := c.rpsService.GetCPMKByRPS(rpsID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPMK",
		Data:    response,
	})
}

// ========== Rencana Pembelajaran Endpoints ==========

// AddRencanaPembelajaran godoc
// @Summary Add Rencana Pembelajaran to RPS
// @Description Add Rencana Pembelajaran (Weekly Plan) to RPS
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Param request body dto.RPSRencanaPembelajaranRequest true "Create Rencana Pembelajaran Request"
// @Success 201 {object} dto.APIResponse{data=dto.RPSRencanaPembelajaranResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/rencana-pembelajaran [post]
func (c *RPSController) AddRencanaPembelajaran(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	var req dto.RPSRencanaPembelajaranRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsService.AddRencanaPembelajaran(rpsID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Rencana Pembelajaran berhasil ditambahkan",
		Data:    response,
	})
}

// UpdateRencanaPembelajaran godoc
// @Summary Update Rencana Pembelajaran
// @Description Update Rencana Pembelajaran data
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rencana_id path string true "Rencana Pembelajaran ID"
// @Param request body dto.RPSRencanaPembelajaranRequest true "Update Rencana Pembelajaran Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSRencanaPembelajaranResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/rencana-pembelajaran/{rencana_id} [put]
func (c *RPSController) UpdateRencanaPembelajaran(ctx *gin.Context) {
	rencanaID := ctx.Param("rencana_id")

	var req dto.RPSRencanaPembelajaranRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsService.UpdateRencanaPembelajaran(rencanaID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Rencana Pembelajaran berhasil diupdate",
		Data:    response,
	})
}

// DeleteRencanaPembelajaran godoc
// @Summary Delete Rencana Pembelajaran
// @Description Delete Rencana Pembelajaran from RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param rencana_id path string true "Rencana Pembelajaran ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/rencana-pembelajaran/{rencana_id} [delete]
func (c *RPSController) DeleteRencanaPembelajaran(ctx *gin.Context) {
	rencanaID := ctx.Param("rencana_id")

	if err := c.rpsService.DeleteRencanaPembelajaran(rencanaID); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Rencana Pembelajaran berhasil dihapus",
	})
}

// GetRencanaPembelajaranByRPS godoc
// @Summary Get Rencana Pembelajaran by RPS ID
// @Description Get all Rencana Pembelajaran for a specific RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSRencanaPembelajaranResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/rencana-pembelajaran [get]
func (c *RPSController) GetRencanaPembelajaranByRPS(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	response, err := c.rpsService.GetRencanaPembelajaranByRPS(rpsID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data Rencana Pembelajaran",
		Data:    response,
	})
}

// ========== Bahan Bacaan Endpoints ==========

// AddBahanBacaan godoc
// @Summary Add Bahan Bacaan to RPS
// @Description Add Bahan Bacaan (Reference) to RPS
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Param request body dto.RPSBahanBacaanRequest true "Create Bahan Bacaan Request"
// @Success 201 {object} dto.APIResponse{data=dto.RPSBahanBacaanResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/bahan-bacaan [post]
func (c *RPSController) AddBahanBacaan(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	var req dto.RPSBahanBacaanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsService.AddBahanBacaan(rpsID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Bahan Bacaan berhasil ditambahkan",
		Data:    response,
	})
}

// UpdateBahanBacaan godoc
// @Summary Update Bahan Bacaan
// @Description Update Bahan Bacaan data
// @Tags RPS
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param bahan_id path string true "Bahan Bacaan ID"
// @Param request body dto.RPSBahanBacaanRequest true "Update Bahan Bacaan Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSBahanBacaanResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/bahan-bacaan/{bahan_id} [put]
func (c *RPSController) UpdateBahanBacaan(ctx *gin.Context) {
	bahanID := ctx.Param("bahan_id")

	var req dto.RPSBahanBacaanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsService.UpdateBahanBacaan(bahanID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Bahan Bacaan berhasil diupdate",
		Data:    response,
	})
}

// DeleteBahanBacaan godoc
// @Summary Delete Bahan Bacaan
// @Description Delete Bahan Bacaan from RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param bahan_id path string true "Bahan Bacaan ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/bahan-bacaan/{bahan_id} [delete]
func (c *RPSController) DeleteBahanBacaan(ctx *gin.Context) {
	bahanID := ctx.Param("bahan_id")

	if err := c.rpsService.DeleteBahanBacaan(bahanID); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Bahan Bacaan berhasil dihapus",
	})
}

// GetBahanBacaanByRPS godoc
// @Summary Get Bahan Bacaan by RPS ID
// @Description Get all Bahan Bacaan for a specific RPS
// @Tags RPS
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSBahanBacaanResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/bahan-bacaan [get]
func (c *RPSController) GetBahanBacaanByRPS(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	response, err := c.rpsService.GetBahanBacaanByRPS(rpsID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data Bahan Bacaan",
		Data:    response,
	})
}