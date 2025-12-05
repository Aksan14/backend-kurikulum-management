package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	dashboardService service.DashboardService
}

func NewDashboardController(dashboardService service.DashboardService) *DashboardController {
	return &DashboardController{dashboardService: dashboardService}
}

// GetKaprodiDashboard godoc
// @Summary Get Kaprodi Dashboard
// @Description Get dashboard statistics for Kaprodi
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse{data=dto.KaprodiDashboardResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /dashboard/kaprodi [get]
func (c *DashboardController) GetKaprodiDashboard(ctx *gin.Context) {
	response, err := c.dashboardService.GetKaprodiDashboard()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data dashboard Kaprodi",
		Data:    response,
	})
}

// GetDosenDashboard godoc
// @Summary Get Dosen Dashboard
// @Description Get dashboard statistics for Dosen
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse{data=dto.DosenDashboardResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /dashboard/dosen [get]
func (c *DashboardController) GetDosenDashboard(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	response, err := c.dashboardService.GetDosenDashboard(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data dashboard Dosen",
		Data:    response,
	})
}

// GetMyDashboard godoc
// @Summary Get My Dashboard
// @Description Get dashboard based on current user's role
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /dashboard [get]
func (c *DashboardController) GetMyDashboard(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	role, _ := ctx.Get("role")

	var response interface{}
	var err error

	if role.(string) == "kaprodi" {
		response, err = c.dashboardService.GetKaprodiDashboard()
	} else {
		response, err = c.dashboardService.GetDosenDashboard(userID.(string))
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data dashboard",
		Data:    response,
	})
}
