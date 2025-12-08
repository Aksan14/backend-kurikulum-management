package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CPLAssignmentController struct {
	assignmentService service.CPLAssignmentService
}

func NewCPLAssignmentController(assignmentService service.CPLAssignmentService) *CPLAssignmentController {
	return &CPLAssignmentController{assignmentService: assignmentService}
}

// GetAllAssignments godoc
// @Summary Get all CPL Assignments
// @Description Get all CPL Assignments dengan pagination
// @Tags CPLAssignment
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param dosen_id query string false "Filter by dosen"
// @Param cpl_id query string false "Filter by CPL"
// @Param status query string false "Filter by status"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /cpl-assignments [get]
func (c *CPLAssignmentController) GetAllAssignments(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := dto.CPLAssignmentListRequest{
		Page:      page,
		Limit:     limit,
		DosenID:   ctx.Query("dosen_id"),
		CPLID:     ctx.Query("cpl_id"),
		Status:    ctx.Query("status"),
		SortBy:    ctx.Query("sort_by"),
		SortOrder: ctx.Query("sort_order"),
	}

	response, err := c.assignmentService.GetAllAssignments(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPL Assignment",
		Data:    response,
	})
}

// GetAssignmentByID godoc
// @Summary Get CPL Assignment by ID
// @Description Get detail CPL Assignment berdasarkan ID
// @Tags CPLAssignment
// @Produce json
// @Security BearerAuth
// @Param id path string true "Assignment ID"
// @Success 200 {object} dto.APIResponse{data=dto.CPLAssignmentResponse}
// @Failure 404 {object} dto.ErrorResponse
// @Router /cpl-assignments/{id} [get]
func (c *CPLAssignmentController) GetAssignmentByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.assignmentService.GetAssignmentByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPL Assignment",
		Data:    response,
	})
}

// CreateAssignment godoc
// @Summary Create new CPL Assignment
// @Description Create new CPL Assignment (Kaprodi only)
// @Tags CPLAssignment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateCPLAssignmentRequest true "Create Assignment Request"
// @Success 201 {object} dto.APIResponse{data=[]dto.CPLAssignmentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /cpl-assignments [post]
func (c *CPLAssignmentController) CreateAssignment(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var req dto.CreateCPLAssignmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	// Validate that either CPL IDs are provided or MataKuliahID is provided
	if len(req.CPLIDs) == 0 && req.MataKuliahID == nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   "harus menyediakan cpl_ids atau mata_kuliah_id",
		})
		return
	}

	responses, err := c.assignmentService.CreateAssignment(userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: fmt.Sprintf("%d CPL Assignment berhasil dibuat", len(responses)),
		Data:    responses,
	})
}

// DeleteAssignment godoc
// @Summary Delete CPL Assignment
// @Description Delete CPL Assignment (Kaprodi only)
// @Tags CPLAssignment
// @Produce json
// @Security BearerAuth
// @Param id path string true "Assignment ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /cpl-assignments/{id} [delete]
func (c *CPLAssignmentController) DeleteAssignment(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.assignmentService.DeleteAssignment(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "CPL Assignment berhasil dihapus",
	})
}

// UpdateAssignmentStatus godoc
// @Summary Update CPL Assignment status
// @Description Update status CPL Assignment
// @Tags CPLAssignment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Assignment ID"
// @Param request body dto.UpdateCPLAssignmentStatusRequest true "Update Status Request"
// @Success 200 {object} dto.APIResponse{data=dto.CPLAssignmentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /cpl-assignments/{id}/status [patch]
func (c *CPLAssignmentController) UpdateAssignmentStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UpdateCPLAssignmentStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.assignmentService.UpdateAssignmentStatus(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Status CPL Assignment berhasil diupdate",
		Data:    response,
	})
}

// GetMyAssignments godoc
// @Summary Get my CPL Assignments
// @Description Get all CPL Assignments for current user (Dosen)
// @Tags CPLAssignment
// @Produce json
// @Security BearerAuth
// @Param status query string false "Filter by status"
// @Success 200 {object} dto.APIResponse{data=[]dto.CPLAssignmentResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /cpl-assignments/my [get]
func (c *CPLAssignmentController) GetMyAssignments(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	status := ctx.Query("status")

	response, err := c.assignmentService.GetAssignmentsByDosen(userID.(string), status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPL Assignment",
		Data:    response,
	})
}

// GetAssignmentsByCPL godoc
// @Summary Get CPL Assignments by CPL
// @Description Get all CPL Assignments for specific CPL
// @Tags CPLAssignment
// @Produce json
// @Security BearerAuth
// @Param cpl_id path string true "CPL ID"
// @Success 200 {object} dto.APIResponse{data=[]dto.CPLAssignmentResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /cpl-assignments/cpl/{cpl_id} [get]
func (c *CPLAssignmentController) GetAssignmentsByCPL(ctx *gin.Context) {
	cplID := ctx.Param("cpl_id")

	response, err := c.assignmentService.GetAssignmentsByCPL(cplID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data CPL Assignment",
		Data:    response,
	})
}
