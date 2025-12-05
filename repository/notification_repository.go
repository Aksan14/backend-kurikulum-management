package repository

import (
	"backend-kurikulum-apps/model"
	"time"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification *model.Notification) error
	CreateBatch(notifications []model.Notification) error
	FindByID(id string) (*model.Notification, error)
	FindByUserID(userID string, page, limit int, notifType string, isRead *bool, sortOrder string) ([]model.Notification, int64, error)
	Update(notification *model.Notification) error
	MarkAsRead(id string) error
	MarkAllAsRead(userID string) error
	Delete(id string) error
	DeleteByUserID(userID string) error
	CountUnread(userID string) (int64, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(notification *model.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) CreateBatch(notifications []model.Notification) error {
	return r.db.Create(&notifications).Error
}

func (r *notificationRepository) FindByID(id string) (*model.Notification, error) {
	var notification model.Notification
	err := r.db.Where("id = ?", id).First(&notification).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *notificationRepository) FindByUserID(userID string, page, limit int, notifType string, isRead *bool, sortOrder string) ([]model.Notification, int64, error) {
	var notifications []model.Notification
	var total int64

	query := r.db.Model(&model.Notification{}).Where("user_id = ?", userID)

	if notifType != "" {
		query = query.Where("type = ?", notifType)
	}
	if isRead != nil {
		query = query.Where("is_read = ?", *isRead)
	}

	query.Count(&total)

	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Order("created_at " + sortOrder).Offset(offset).Limit(limit).Find(&notifications).Error
	return notifications, total, err
}

func (r *notificationRepository) Update(notification *model.Notification) error {
	return r.db.Save(notification).Error
}

func (r *notificationRepository) MarkAsRead(id string) error {
	return r.db.Model(&model.Notification{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_read": true,
		"read_at": time.Now(),
	}).Error
}

func (r *notificationRepository) MarkAllAsRead(userID string) error {
	return r.db.Model(&model.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Updates(map[string]interface{}{
		"is_read": true,
		"read_at": time.Now(),
	}).Error
}

func (r *notificationRepository) Delete(id string) error {
	return r.db.Delete(&model.Notification{}, "id = ?", id).Error
}

func (r *notificationRepository) DeleteByUserID(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Notification{}).Error
}

func (r *notificationRepository) CountUnread(userID string) (int64, error) {
	var count int64
	err := r.db.Model(&model.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&count).Error
	return count, err
}
