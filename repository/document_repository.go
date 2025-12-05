package repository

import (
	"backend-kurikulum-apps/model"

	"gorm.io/gorm"
)

// Document Template Repository
type DocumentTemplateRepository interface {
	Create(template *model.DocumentTemplate) error
	FindByID(id string) (*model.DocumentTemplate, error)
	FindAll(page, limit int, search string, isActive *bool, sortBy, sortOrder string) ([]model.DocumentTemplate, int64, error)
	Update(template *model.DocumentTemplate) error
	Delete(id string) error
}

type documentTemplateRepository struct {
	db *gorm.DB
}

func NewDocumentTemplateRepository(db *gorm.DB) DocumentTemplateRepository {
	return &documentTemplateRepository{db: db}
}

func (r *documentTemplateRepository) Create(template *model.DocumentTemplate) error {
	return r.db.Create(template).Error
}

func (r *documentTemplateRepository) FindByID(id string) (*model.DocumentTemplate, error) {
	var template model.DocumentTemplate
	err := r.db.Preload("Creator").Where("id = ?", id).First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *documentTemplateRepository) FindAll(page, limit int, search string, isActive *bool, sortBy, sortOrder string) ([]model.DocumentTemplate, int64, error) {
	var templates []model.DocumentTemplate
	var total int64

	query := r.db.Model(&model.DocumentTemplate{})

	if search != "" {
		query = query.Where("nama LIKE ? OR deskripsi LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if isActive != nil {
		query = query.Where("is_active = ?", *isActive)
	}

	query.Count(&total)

	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Preload("Creator").Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&templates).Error
	return templates, total, err
}

func (r *documentTemplateRepository) Update(template *model.DocumentTemplate) error {
	return r.db.Save(template).Error
}

func (r *documentTemplateRepository) Delete(id string) error {
	return r.db.Delete(&model.DocumentTemplate{}, "id = ?", id).Error
}

// Generated Document Repository
type GeneratedDocumentRepository interface {
	Create(doc *model.GeneratedDocument) error
	FindByID(id string) (*model.GeneratedDocument, error)
	FindAll(page, limit int, templateID, status, tahun, fileType, sortBy, sortOrder string) ([]model.GeneratedDocument, int64, error)
	Update(doc *model.GeneratedDocument) error
	Delete(id string) error
	UpdateProgress(id string, progress int) error
	UpdateStatus(id, status string, fileURL *string, fileSize *int64, errorMessage *string) error
	CountByStatus(status string) (int64, error)
}

type generatedDocumentRepository struct {
	db *gorm.DB
}

func NewGeneratedDocumentRepository(db *gorm.DB) GeneratedDocumentRepository {
	return &generatedDocumentRepository{db: db}
}

func (r *generatedDocumentRepository) Create(doc *model.GeneratedDocument) error {
	return r.db.Create(doc).Error
}

func (r *generatedDocumentRepository) FindByID(id string) (*model.GeneratedDocument, error) {
	var doc model.GeneratedDocument
	err := r.db.Preload("Template").Preload("Creator").Where("id = ?", id).First(&doc).Error
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *generatedDocumentRepository) FindAll(page, limit int, templateID, status, tahun, fileType, sortBy, sortOrder string) ([]model.GeneratedDocument, int64, error) {
	var docs []model.GeneratedDocument
	var total int64

	query := r.db.Model(&model.GeneratedDocument{})

	if templateID != "" {
		query = query.Where("template_id = ?", templateID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if tahun != "" {
		query = query.Where("tahun = ?", tahun)
	}
	if fileType != "" {
		query = query.Where("file_type = ?", fileType)
	}

	query.Count(&total)

	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Preload("Template").Preload("Creator").Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&docs).Error
	return docs, total, err
}

func (r *generatedDocumentRepository) Update(doc *model.GeneratedDocument) error {
	return r.db.Save(doc).Error
}

func (r *generatedDocumentRepository) Delete(id string) error {
	return r.db.Delete(&model.GeneratedDocument{}, "id = ?", id).Error
}

func (r *generatedDocumentRepository) UpdateProgress(id string, progress int) error {
	return r.db.Model(&model.GeneratedDocument{}).Where("id = ?", id).Update("progress", progress).Error
}

func (r *generatedDocumentRepository) UpdateStatus(id, status string, fileURL *string, fileSize *int64, errorMessage *string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == "ready" {
		updates["progress"] = 100
		updates["completed_at"] = gorm.Expr("NOW()")
		if fileURL != nil {
			updates["file_url"] = *fileURL
		}
		if fileSize != nil {
			updates["file_size"] = *fileSize
		}
	} else if status == "failed" {
		if errorMessage != nil {
			updates["error_message"] = *errorMessage
		}
	}
	return r.db.Model(&model.GeneratedDocument{}).Where("id = ?", id).Updates(updates).Error
}

func (r *generatedDocumentRepository) CountByStatus(status string) (int64, error) {
	var count int64
	query := r.db.Model(&model.GeneratedDocument{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&count).Error
	return count, err
}
