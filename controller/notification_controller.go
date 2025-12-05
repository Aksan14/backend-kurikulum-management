package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationService service.NotificationService
}

func NewNotificationController(notificationService service.NotificationService) *NotificationController {
	return &NotificationController{notificationService: notificationService}
}

// GetMyNotifications godoc
// @Summary Get my notifications
// @Description Get all notifications for current user
// @Tags Notification
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param is_read query bool false "Filter by read status"
// @Param type query string false "Filter by type"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /notifications [get]
func (c *NotificationController) GetMyNotifications(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := dto.NotificationListRequest{
		Page:   page,
		Limit:  limit,
		UserID: userID.(string),
		Type:   ctx.Query("type"),
	}

	if isRead := ctx.Query("is_read"); isRead != "" {
		val, _ := strconv.ParseBool(isRead)
		req.IsRead = &val
	}

	response, err := c.notificationService.GetNotifications(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan notifikasi",
		Data:    response,
	})
}

// GetUnreadCount godoc
// @Summary Get unread notification count
// @Description Get count of unread notifications
// @Tags Notification
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /notifications/unread-count [get]
func (c *NotificationController) GetUnreadCount(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	count, err := c.notificationService.GetUnreadCount(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan jumlah notifikasi belum dibaca",
		Data:    map[string]int64{"unread_count": count},
	})
}

// MarkAsRead godoc
// @Summary Mark notification as read
// @Description Mark specific notification as read
// @Tags Notification
// @Produce json
// @Security BearerAuth
// @Param id path string true "Notification ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /notifications/{id}/read [patch]
func (c *NotificationController) MarkAsRead(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.notificationService.MarkAsRead(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Notifikasi berhasil ditandai sebagai dibaca",
	})
}

// MarkAllAsRead godoc
// @Summary Mark all notifications as read
// @Description Mark all notifications for current user as read
// @Tags Notification
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /notifications/read-all [patch]
func (c *NotificationController) MarkAllAsRead(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	if err := c.notificationService.MarkAllAsRead(userID.(string)); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Semua notifikasi berhasil ditandai sebagai dibaca",
	})
}

// DeleteNotification godoc
// @Summary Delete notification
// @Description Delete specific notification
// @Tags Notification
// @Produce json
// @Security BearerAuth
// @Param id path string true "Notification ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /notifications/{id} [delete]
func (c *NotificationController) DeleteNotification(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.notificationService.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Notifikasi berhasil dihapus",
	})
}

// CreateNotification godoc
// @Summary Create notification (Admin only)
// @Description Create a new notification
// @Tags Notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateNotificationRequest true "Create Notification Request"
// @Success 201 {object} dto.APIResponse{data=dto.NotificationResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /notifications [post]
func (c *NotificationController) CreateNotification(ctx *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.notificationService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Notifikasi berhasil dibuat",
		Data:    response,
	})
}
