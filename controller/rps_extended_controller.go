package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RPSExtendedController struct {
	rpsExtendedService service.RPSExtendedService
}

func NewRPSExtendedController(rpsExtendedService service.RPSExtendedService) *RPSExtendedController {
	return &RPSExtendedController{rpsExtendedService: rpsExtendedService}
}

// ============ SUB CPMK ENDPOINTS ============

// AddSubCPMK godoc
// @Summary Add Sub-CPMK to CPMK
// @Description Add new Sub-CPMK to existing CPMK
// @Tags RPS-SubCPMK
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cpmk_id path string true "CPMK ID"
// @Param request body dto.SubCPMKRequest true "Sub-CPMK Request"
// @Success 201 {object} dto.APIResponse{data=dto.SubCPMKResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/cpmk/{cpmk_id}/sub-cpmk [post]
func (c *RPSExtendedController) AddSubCPMK(ctx *gin.Context) {
	cpmkID := ctx.Param("cpmk_id")

	var req dto.SubCPMKRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.AddSubCPMK(cpmkID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Sub-CPMK berhasil ditambahkan",
		Data:    response,
	})
}

// GetSubCPMKByCPMK godoc
// @Summary Get all Sub-CPMK by CPMK ID
// @Description Get all Sub-CPMK for a specific CPMK
// @Tags RPS-SubCPMK
// @Produce json
// @Security BearerAuth
// @Param cpmk_id path string true "CPMK ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.SubCPMKResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/cpmk/{cpmk_id}/sub-cpmk [get]
func (c *RPSExtendedController) GetSubCPMKByCPMK(ctx *gin.Context) {
	cpmkID := ctx.Param("cpmk_id")

	response, err := c.rpsExtendedService.GetSubCPMKByCPMK(cpmkID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data Sub-CPMK",
		Data:    response,
	})
}

// UpdateSubCPMK godoc
// @Summary Update Sub-CPMK
// @Description Update existing Sub-CPMK data
// @Tags RPS-SubCPMK
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param sub_cpmk_id path string true "Sub-CPMK ID"
// @Param request body dto.SubCPMKRequest true "Update Sub-CPMK Request"
// @Success 200 {object} dto.APIResponse{data=dto.SubCPMKResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/sub-cpmk/{sub_cpmk_id} [put]
func (c *RPSExtendedController) UpdateSubCPMK(ctx *gin.Context) {
	subCPMKID := ctx.Param("sub_cpmk_id")

	var req dto.SubCPMKRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.UpdateSubCPMK(subCPMKID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Sub-CPMK berhasil diupdate",
		Data:    response,
	})
}

// DeleteSubCPMK godoc
// @Summary Delete Sub-CPMK
// @Description Delete Sub-CPMK from CPMK
// @Tags RPS-SubCPMK
// @Produce json
// @Security BearerAuth
// @Param sub_cpmk_id path string true "Sub-CPMK ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/sub-cpmk/{sub_cpmk_id} [delete]
func (c *RPSExtendedController) DeleteSubCPMK(ctx *gin.Context) {
	subCPMKID := ctx.Param("sub_cpmk_id")

	if err := c.rpsExtendedService.DeleteSubCPMK(subCPMKID); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Sub-CPMK berhasil dihapus",
	})
}

// ============ RENCANA TUGAS ENDPOINTS ============

// AddRencanaTugas godoc
// @Summary Add Rencana Tugas to RPS
// @Description Add new Rencana Tugas to existing RPS
// @Tags RPS-RencanaTugas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Param request body dto.RPSRencanaTugasRequest true "Rencana Tugas Request"
// @Success 201 {object} dto.APIResponse{data=dto.RPSRencanaTugasResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/rencana-tugas [post]
func (c *RPSExtendedController) AddRencanaTugas(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	var req dto.RPSRencanaTugasRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.AddRencanaTugas(rpsID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Rencana Tugas berhasil ditambahkan",
		Data:    response,
	})
}

// GetRencanaTugasByRPS godoc
// @Summary Get all Rencana Tugas by RPS ID
// @Description Get all Rencana Tugas for a specific RPS
// @Tags RPS-RencanaTugas
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSRencanaTugasResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/rencana-tugas [get]
func (c *RPSExtendedController) GetRencanaTugasByRPS(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	response, err := c.rpsExtendedService.GetRencanaTugasByRPS(rpsID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data Rencana Tugas",
		Data:    response,
	})
}

// UpdateRencanaTugas godoc
// @Summary Update Rencana Tugas
// @Description Update existing Rencana Tugas data
// @Tags RPS-RencanaTugas
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param tugas_id path string true "Rencana Tugas ID"
// @Param request body dto.RPSRencanaTugasRequest true "Update Rencana Tugas Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSRencanaTugasResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/rencana-tugas/{tugas_id} [put]
func (c *RPSExtendedController) UpdateRencanaTugas(ctx *gin.Context) {
	tugasID := ctx.Param("tugas_id")

	var req dto.RPSRencanaTugasRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.UpdateRencanaTugas(tugasID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Rencana Tugas berhasil diupdate",
		Data:    response,
	})
}

// DeleteRencanaTugas godoc
// @Summary Delete Rencana Tugas
// @Description Delete Rencana Tugas from RPS
// @Tags RPS-RencanaTugas
// @Produce json
// @Security BearerAuth
// @Param tugas_id path string true "Rencana Tugas ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/rencana-tugas/{tugas_id} [delete]
func (c *RPSExtendedController) DeleteRencanaTugas(ctx *gin.Context) {
	tugasID := ctx.Param("tugas_id")

	if err := c.rpsExtendedService.DeleteRencanaTugas(tugasID); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Rencana Tugas berhasil dihapus",
	})
}

// ============ ANALISIS KETERCAPAIAN CPL ENDPOINTS ============

// AddAnalisisKetercapaianCPL godoc
// @Summary Add Analisis Ketercapaian CPL to RPS
// @Description Add new Analisis Ketercapaian CPL to existing RPS
// @Tags RPS-AnalisisKetercapaianCPL
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Param request body dto.RPSAnalisisKetercapaianCPLRequest true "Analisis Ketercapaian CPL Request"
// @Success 201 {object} dto.APIResponse{data=dto.RPSAnalisisKetercapaianCPLResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/analisis-ketercapaian [post]
func (c *RPSExtendedController) AddAnalisisKetercapaianCPL(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	var req dto.RPSAnalisisKetercapaianCPLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.AddAnalisisKetercapaianCPL(rpsID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Analisis Ketercapaian CPL berhasil ditambahkan",
		Data:    response,
	})
}

// GetAnalisisKetercapaianCPLByRPS godoc
// @Summary Get all Analisis Ketercapaian CPL by RPS ID
// @Description Get all Analisis Ketercapaian CPL for a specific RPS
// @Tags RPS-AnalisisKetercapaianCPL
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSAnalisisKetercapaianCPLResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/analisis-ketercapaian [get]
func (c *RPSExtendedController) GetAnalisisKetercapaianCPLByRPS(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	response, err := c.rpsExtendedService.GetAnalisisKetercapaianCPLByRPS(rpsID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data Analisis Ketercapaian CPL",
		Data:    response,
	})
}

// UpdateAnalisisKetercapaianCPL godoc
// @Summary Update Analisis Ketercapaian CPL
// @Description Update existing Analisis Ketercapaian CPL data
// @Tags RPS-AnalisisKetercapaianCPL
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param analisis_id path string true "Analisis Ketercapaian CPL ID"
// @Param request body dto.RPSAnalisisKetercapaianCPLRequest true "Update Analisis Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSAnalisisKetercapaianCPLResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/analisis-ketercapaian/{analisis_id} [put]
func (c *RPSExtendedController) UpdateAnalisisKetercapaianCPL(ctx *gin.Context) {
	analisisID := ctx.Param("analisis_id")

	var req dto.RPSAnalisisKetercapaianCPLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.UpdateAnalisisKetercapaianCPL(analisisID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Analisis Ketercapaian CPL berhasil diupdate",
		Data:    response,
	})
}

// DeleteAnalisisKetercapaianCPL godoc
// @Summary Delete Analisis Ketercapaian CPL
// @Description Delete Analisis Ketercapaian CPL from RPS
// @Tags RPS-AnalisisKetercapaianCPL
// @Produce json
// @Security BearerAuth
// @Param analisis_id path string true "Analisis Ketercapaian CPL ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/analisis-ketercapaian/{analisis_id} [delete]
func (c *RPSExtendedController) DeleteAnalisisKetercapaianCPL(ctx *gin.Context) {
	analisisID := ctx.Param("analisis_id")

	if err := c.rpsExtendedService.DeleteAnalisisKetercapaianCPL(analisisID); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Analisis Ketercapaian CPL berhasil dihapus",
	})
}

// ============ SKALA PENILAIAN ENDPOINTS ============

// AddSkalaPenilaian godoc
// @Summary Add Skala Penilaian to RPS
// @Description Add new Skala Penilaian to existing RPS
// @Tags RPS-SkalaPenilaian
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Param request body dto.RPSSkalaPenilaianRequest true "Skala Penilaian Request"
// @Success 201 {object} dto.APIResponse{data=dto.RPSSkalaPenilaianResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/skala-penilaian [post]
func (c *RPSExtendedController) AddSkalaPenilaian(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	var req dto.RPSSkalaPenilaianRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.AddSkalaPenilaian(rpsID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Skala Penilaian berhasil ditambahkan",
		Data:    response,
	})
}

// GetSkalaPenilaianByRPS godoc
// @Summary Get all Skala Penilaian by RPS ID
// @Description Get all Skala Penilaian for a specific RPS
// @Tags RPS-SkalaPenilaian
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSSkalaPenilaianResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/skala-penilaian [get]
func (c *RPSExtendedController) GetSkalaPenilaianByRPS(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	response, err := c.rpsExtendedService.GetSkalaPenilaianByRPS(rpsID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data Skala Penilaian",
		Data:    response,
	})
}

// UpdateSkalaPenilaian godoc
// @Summary Update Skala Penilaian
// @Description Update existing Skala Penilaian data
// @Tags RPS-SkalaPenilaian
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param skala_id path string true "Skala Penilaian ID"
// @Param request body dto.RPSSkalaPenilaianRequest true "Update Skala Penilaian Request"
// @Success 200 {object} dto.APIResponse{data=dto.RPSSkalaPenilaianResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/skala-penilaian/{skala_id} [put]
func (c *RPSExtendedController) UpdateSkalaPenilaian(ctx *gin.Context) {
	skalaID := ctx.Param("skala_id")

	var req dto.RPSSkalaPenilaianRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.UpdateSkalaPenilaian(skalaID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Skala Penilaian berhasil diupdate",
		Data:    response,
	})
}

// DeleteSkalaPenilaian godoc
// @Summary Delete Skala Penilaian
// @Description Delete Skala Penilaian from RPS
// @Tags RPS-SkalaPenilaian
// @Produce json
// @Security BearerAuth
// @Param skala_id path string true "Skala Penilaian ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/skala-penilaian/{skala_id} [delete]
func (c *RPSExtendedController) DeleteSkalaPenilaian(ctx *gin.Context) {
	skalaID := ctx.Param("skala_id")

	if err := c.rpsExtendedService.DeleteSkalaPenilaian(skalaID); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Skala Penilaian berhasil dihapus",
	})
}

// SetDefaultSkalaPenilaian godoc
// @Summary Set Default Skala Penilaian
// @Description Set default Skala Penilaian values for RPS
// @Tags RPS-SkalaPenilaian
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rps_id path string true "RPS ID"
// @Param request body dto.RPSBatchSkalaPenilaianRequest true "Batch Skala Penilaian Request"
// @Success 200 {object} dto.APIResponse{data=[]dto.RPSSkalaPenilaianResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /rps/{rps_id}/skala-penilaian/batch [post]
func (c *RPSExtendedController) SetDefaultSkalaPenilaian(ctx *gin.Context) {
	rpsID := ctx.Param("rps_id")

	var req dto.RPSBatchSkalaPenilaianRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.rpsExtendedService.SetDefaultSkalaPenilaian(rpsID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Skala Penilaian berhasil diset",
		Data:    response,
	})
}
