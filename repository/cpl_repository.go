package repository

import (
	"backend-kurikulum-apps/model"
	"time"

	"gorm.io/gorm"
)

type CPLRepository interface {
	Create(cpl *model.CPL) error
	FindByID(id string) (*model.CPL, error)
	FindByKode(kode string) (*model.CPL, error)
	FindAll(page, limit int, search, status, aspek, kategori, sortBy, sortOrder string) ([]model.CPL, int64, error)
	Update(cpl *model.CPL) error
	Delete(id string) error
	UpdateStatus(id, status string) error
	Count() (int64, error)
	CountByStatus(status string) (int64, error)
	FindByStatus(status string) ([]model.CPL, error)
}

type cplRepository struct {
	db *gorm.DB
}

func NewCPLRepository(db *gorm.DB) CPLRepository {
	return &cplRepository{db: db}
}

func (r *cplRepository) Create(cpl *model.CPL) error {
	return r.db.Create(cpl).Error
}

func (r *cplRepository) FindByID(id string) (*model.CPL, error) {
	var cpl model.CPL
	err := r.db.Preload("Creator").Where("id = ?", id).First(&cpl).Error
	if err != nil {
		return nil, err
	}
	return &cpl, nil
}

func (r *cplRepository) FindByKode(kode string) (*model.CPL, error) {
	var cpl model.CPL
	err := r.db.Where("kode = ?", kode).First(&cpl).Error
	if err != nil {
		return nil, err
	}
	return &cpl, nil
}

func (r *cplRepository) FindAll(page, limit int, search, status, aspek, kategori, sortBy, sortOrder string) ([]model.CPL, int64, error) {
	var cpls []model.CPL
	var total int64

	query := r.db.Model(&model.CPL{})

	if search != "" {
		query = query.Where("kode LIKE ? OR judul LIKE ? OR deskripsi LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if aspek != "" {
		query = query.Where("aspek = ?", aspek)
	}
	if kategori != "" {
		query = query.Where("kategori = ?", kategori)
	}

	query.Count(&total)

	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Preload("Creator").Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&cpls).Error
	return cpls, total, err
}

func (r *cplRepository) Update(cpl *model.CPL) error {
	return r.db.Save(cpl).Error
}

func (r *cplRepository) Delete(id string) error {
	return r.db.Delete(&model.CPL{}, "id = ?", id).Error
}

func (r *cplRepository) UpdateStatus(id, status string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if status == "published" {
		updates["published_at"] = time.Now()
	} else if status == "archived" {
		updates["archived_at"] = time.Now()
	}
	return r.db.Model(&model.CPL{}).Where("id = ?", id).Updates(updates).Error
}

func (r *cplRepository) CountByStatus(status string) (int64, error) {
	var count int64
	query := r.db.Model(&model.CPL{})
	if status != "" {
		query = query.Where("status = ?", status)
	} else {
		query = query.Where("status != ?", "archived")
	}
	err := query.Count(&count).Error
	return count, err
}

func (r *cplRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&model.CPL{}).Count(&count).Error
	return count, err
}

func (r *cplRepository) FindByStatus(status string) ([]model.CPL, error) {
	var cpls []model.CPL
	err := r.db.Preload("Creator").Where("status = ?", status).Order("kode ASC").Find(&cpls).Error
	return cpls, err
}
