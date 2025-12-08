package repository

import (
	"backend-kurikulum-apps/model"

	"gorm.io/gorm"
)

type CPLMKMappingRepository interface {
	Create(mapping *model.CPLMKMapping) error
	FindByID(id string) (*model.CPLMKMapping, error)
	FindByCPLAndMK(cplID, mataKuliahID string) (*model.CPLMKMapping, error)
	FindByMataKuliahID(mataKuliahID string) ([]model.CPLMKMapping, error)
	FindAll(page, limit int, cplID, mataKuliahID, level string) ([]model.CPLMKMapping, int64, error)
	Update(mapping *model.CPLMKMapping) error
	Delete(id string) error
}

type cplMKMappingRepository struct {
	db *gorm.DB
}

func NewCPLMKMappingRepository(db *gorm.DB) CPLMKMappingRepository {
	return &cplMKMappingRepository{db: db}
}

func (r *cplMKMappingRepository) Create(mapping *model.CPLMKMapping) error {
	return r.db.Create(mapping).Error
}

func (r *cplMKMappingRepository) FindByID(id string) (*model.CPLMKMapping, error) {
	var mapping model.CPLMKMapping
	err := r.db.Preload("CPL").Preload("MataKuliah").
		Where("id = ?", id).First(&mapping).Error
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

func (r *cplMKMappingRepository) FindByCPLAndMK(cplID, mataKuliahID string) (*model.CPLMKMapping, error) {
	var mapping model.CPLMKMapping
	err := r.db.Preload("CPL").Preload("MataKuliah").
		Where("cpl_id = ? AND mata_kuliah_id = ?", cplID, mataKuliahID).First(&mapping).Error
	if err != nil {
		return nil, err
	}
	return &mapping, nil
}

func (r *cplMKMappingRepository) FindByMataKuliahID(mataKuliahID string) ([]model.CPLMKMapping, error) {
	var mappings []model.CPLMKMapping
	err := r.db.Preload("CPL").Preload("MataKuliah").
		Where("mata_kuliah_id = ?", mataKuliahID).Find(&mappings).Error
	return mappings, err
}

func (r *cplMKMappingRepository) FindAll(page, limit int, cplID, mataKuliahID, level string) ([]model.CPLMKMapping, int64, error) {
	var mappings []model.CPLMKMapping
	var total int64

	query := r.db.Model(&model.CPLMKMapping{})

	if cplID != "" {
		query = query.Where("cpl_id = ?", cplID)
	}
	if mataKuliahID != "" {
		query = query.Where("mata_kuliah_id = ?", mataKuliahID)
	}
	if level != "" {
		query = query.Where("level = ?", level)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("CPL").Preload("MataKuliah").
		Order("created_at DESC").Offset(offset).Limit(limit).Find(&mappings).Error

	return mappings, total, err
}

func (r *cplMKMappingRepository) Update(mapping *model.CPLMKMapping) error {
	return r.db.Save(mapping).Error
}

func (r *cplMKMappingRepository) Delete(id string) error {
	return r.db.Delete(&model.CPLMKMapping{}, "id = ?", id).Error
}
