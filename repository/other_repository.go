package repository

import (
	"backend-kurikulum-apps/model"

	"gorm.io/gorm"
)

// Audit Log Repository
type AuditLogRepository interface {
	Create(log *model.AuditLog) error
	FindByID(id string) (*model.AuditLog, error)
	FindAll(page, limit int, userID, action, tableName, startDate, endDate, sortOrder string) ([]model.AuditLog, int64, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) Create(log *model.AuditLog) error {
	return r.db.Create(log).Error
}

func (r *auditLogRepository) FindByID(id string) (*model.AuditLog, error) {
	var log model.AuditLog
	err := r.db.Preload("User").Where("id = ?", id).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *auditLogRepository) FindAll(page, limit int, userID, action, tableName, startDate, endDate, sortOrder string) ([]model.AuditLog, int64, error) {
	var logs []model.AuditLog
	var total int64

	query := r.db.Model(&model.AuditLog{})

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if action != "" {
		query = query.Where("action = ?", action)
	}
	if tableName != "" {
		query = query.Where("table_name = ?", tableName)
	}
	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	query.Count(&total)

	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Preload("User").Order("created_at " + sortOrder).Offset(offset).Limit(limit).Find(&logs).Error
	return logs, total, err
}

// System Setting Repository
type SystemSettingRepository interface {
	Create(setting *model.SystemSetting) error
	FindByKey(key string) (*model.SystemSetting, error)
	FindAll() ([]model.SystemSetting, error)
	FindByCategory(category string) ([]model.SystemSetting, error)
	FindPublic() ([]model.SystemSetting, error)
	Update(setting *model.SystemSetting) error
	Delete(key string) error
}

type systemSettingRepository struct {
	db *gorm.DB
}

func NewSystemSettingRepository(db *gorm.DB) SystemSettingRepository {
	return &systemSettingRepository{db: db}
}

func (r *systemSettingRepository) Create(setting *model.SystemSetting) error {
	return r.db.Create(setting).Error
}

func (r *systemSettingRepository) FindByKey(key string) (*model.SystemSetting, error) {
	var setting model.SystemSetting
	err := r.db.Where("`key` = ?", key).First(&setting).Error
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r *systemSettingRepository) FindAll() ([]model.SystemSetting, error) {
	var settings []model.SystemSetting
	err := r.db.Find(&settings).Error
	return settings, err
}

func (r *systemSettingRepository) FindByCategory(category string) ([]model.SystemSetting, error) {
	var settings []model.SystemSetting
	err := r.db.Where("category = ?", category).Find(&settings).Error
	return settings, err
}

func (r *systemSettingRepository) FindPublic() ([]model.SystemSetting, error) {
	var settings []model.SystemSetting
	err := r.db.Where("is_public = ?", true).Find(&settings).Error
	return settings, err
}

func (r *systemSettingRepository) Update(setting *model.SystemSetting) error {
	return r.db.Save(setting).Error
}

func (r *systemSettingRepository) Delete(key string) error {
	return r.db.Where("`key` = ?", key).Delete(&model.SystemSetting{}).Error
}

// File Repository
type FileRepository interface {
	Create(file *model.File) error
	FindByID(id string) (*model.File, error)
	FindByRelated(relatedID, relatedType string) ([]model.File, error)
	Update(file *model.File) error
	Delete(id string) error
	DeleteTemporary() error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) Create(file *model.File) error {
	return r.db.Create(file).Error
}

func (r *fileRepository) FindByID(id string) (*model.File, error) {
	var file model.File
	err := r.db.Preload("Uploader").Where("id = ?", id).First(&file).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) FindByRelated(relatedID, relatedType string) ([]model.File, error) {
	var files []model.File
	err := r.db.Where("related_id = ? AND related_type = ?", relatedID, relatedType).Find(&files).Error
	return files, err
}

func (r *fileRepository) Update(file *model.File) error {
	return r.db.Save(file).Error
}

func (r *fileRepository) Delete(id string) error {
	return r.db.Delete(&model.File{}, "id = ?", id).Error
}

func (r *fileRepository) DeleteTemporary() error {
	return r.db.Where("is_temporary = ? AND created_at < DATE_SUB(NOW(), INTERVAL 24 HOUR)", true).Delete(&model.File{}).Error
}

// RPS CPL Mapping Repository
type RPSCPLMappingRepository interface {
	Create(mapping *model.RPSCPLMapping) error
	CreateBatch(mappings []model.RPSCPLMapping) error
	FindByRPSID(rpsID string) ([]model.RPSCPLMapping, error)
	FindByCPLID(cplID string) ([]model.RPSCPLMapping, error)
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
}

type rpsCPLMappingRepository struct {
	db *gorm.DB
}

func NewRPSCPLMappingRepository(db *gorm.DB) RPSCPLMappingRepository {
	return &rpsCPLMappingRepository{db: db}
}

func (r *rpsCPLMappingRepository) Create(mapping *model.RPSCPLMapping) error {
	return r.db.Create(mapping).Error
}

func (r *rpsCPLMappingRepository) CreateBatch(mappings []model.RPSCPLMapping) error {
	return r.db.Create(&mappings).Error
}

func (r *rpsCPLMappingRepository) FindByRPSID(rpsID string) ([]model.RPSCPLMapping, error) {
	var mappings []model.RPSCPLMapping
	err := r.db.Preload("CPL").Preload("CPMK").Where("rps_id = ?", rpsID).Find(&mappings).Error
	return mappings, err
}

func (r *rpsCPLMappingRepository) FindByCPLID(cplID string) ([]model.RPSCPLMapping, error) {
	var mappings []model.RPSCPLMapping
	err := r.db.Preload("RPS").Preload("CPMK").Where("cpl_id = ?", cplID).Find(&mappings).Error
	return mappings, err
}

func (r *rpsCPLMappingRepository) Delete(id string) error {
	return r.db.Delete(&model.RPSCPLMapping{}, "id = ?", id).Error
}

func (r *rpsCPLMappingRepository) DeleteByRPSID(rpsID string) error {
	return r.db.Where("rps_id = ?", rpsID).Delete(&model.RPSCPLMapping{}).Error
}
