package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"errors"
	"time"
)

type NotificationService interface {
	GetByUserID(userID string, req dto.NotificationListRequest) (*dto.PaginatedResponse, error)
	GetNotifications(req dto.NotificationListRequest) (*dto.PaginatedResponse, error)
	GetUnreadCount(userID string) (int64, error)
	Create(req dto.CreateNotificationRequest) (*dto.NotificationResponse, error)
	MarkAsRead(id string) error
	MarkAllAsRead(userID string) error
	Delete(id string) error
}

type notificationService struct {
	notifRepo repository.NotificationRepository
}

func NewNotificationService(notifRepo repository.NotificationRepository) NotificationService {
	return &notificationService{notifRepo: notifRepo}
}

func (s *notificationService) GetByUserID(userID string, req dto.NotificationListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 20
	}

	notifications, total, err := s.notifRepo.FindByUserID(
		userID, req.Page, req.Limit, req.Type, req.IsRead, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.NotificationResponse
	for _, n := range notifications {
		responses = append(responses, toNotificationResponse(&n))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *notificationService) GetNotifications(req dto.NotificationListRequest) (*dto.PaginatedResponse, error) {
	return s.GetByUserID(req.UserID, req)
}

func (s *notificationService) GetUnreadCount(userID string) (int64, error) {
	return s.notifRepo.CountUnread(userID)
}

func (s *notificationService) Create(req dto.CreateNotificationRequest) (*dto.NotificationResponse, error) {
	notification := &model.Notification{
		UserID:        req.UserID,
		Title:         req.Title,
		Message:       req.Message,
		Type:          req.Type,
		ActionURL:     req.ActionURL,
		ReferenceID:   req.RelatedID,
		ReferenceType: req.RelatedType,
		IsRead:        false,
	}

	if err := s.notifRepo.Create(notification); err != nil {
		return nil, errors.New("gagal membuat notifikasi")
	}

	resp := toNotificationResponse(notification)
	return &resp, nil
}

func (s *notificationService) MarkAsRead(id string) error {
	_, err := s.notifRepo.FindByID(id)
	if err != nil {
		return errors.New("notifikasi tidak ditemukan")
	}

	return s.notifRepo.MarkAsRead(id)
}

func (s *notificationService) MarkAllAsRead(userID string) error {
	return s.notifRepo.MarkAllAsRead(userID)
}

func (s *notificationService) Delete(id string) error {
	_, err := s.notifRepo.FindByID(id)
	if err != nil {
		return errors.New("notifikasi tidak ditemukan")
	}

	return s.notifRepo.Delete(id)
}

func toNotificationResponse(n *model.Notification) dto.NotificationResponse {
	return dto.NotificationResponse{
		ID:          n.ID,
		UserID:      n.UserID,
		Title:       n.Title,
		Message:     n.Message,
		Type:        n.Type,
		IsRead:      n.IsRead,
		ActionURL:   n.ActionURL,
		RelatedID:   n.ReferenceID,
		RelatedType: n.ReferenceType,
		Priority:    "normal",
		CreatedAt:   n.CreatedAt,
		ReadAt:      n.ReadAt,
	}
}

// Utility function to send notifications for various events
func (s *notificationService) SendAssignmentNotification(dosenID, title, message, relatedID string) {
	s.Create(dto.CreateNotificationRequest{
		UserID:      dosenID,
		Title:       title,
		Message:     message,
		Type:        "assignment",
		RelatedID:   &relatedID,
		RelatedType: ptrString("cpl_assignment"),
		Priority:    "high",
	})
}

func (s *notificationService) SendApprovalNotification(userID, title, message, relatedID string) {
	s.Create(dto.CreateNotificationRequest{
		UserID:      userID,
		Title:       title,
		Message:     message,
		Type:        "approval",
		RelatedID:   &relatedID,
		RelatedType: ptrString("rps"),
		Priority:    "normal",
	})
}

func (s *notificationService) SendRejectionNotification(userID, title, message, relatedID string) {
	s.Create(dto.CreateNotificationRequest{
		UserID:      userID,
		Title:       title,
		Message:     message,
		Type:        "rejection",
		RelatedID:   &relatedID,
		RelatedType: ptrString("rps"),
		Priority:    "high",
	})
}

func (s *notificationService) SendDeadlineNotification(userID, title, message, relatedID string, deadline time.Time) {
	notification := &model.Notification{
		UserID:        userID,
		Title:         title,
		Message:       message,
		Type:          "deadline",
		ReferenceID:   &relatedID,
		ReferenceType: ptrString("cpl_assignment"),
		IsRead:        false,
	}

	s.notifRepo.Create(notification)
}

func ptrString(s string) *string {
	return &s
}
