package repository

import (
	"backend-kurikulum-apps/model"
	"time"

	"gorm.io/gorm"
)

type RPSRepository interface {
	Create(rps *model.RPS) error
	FindByID(id string) (*model.RPS, error)
	FindAll(page, limit int, search, mataKuliahID, dosenID, status, tahunAkademik string, semester int, sortBy, sortOrder string) ([]model.RPS, int64, error)
	FindByDosenID(dosenID string, page, limit int, status string) ([]model.RPS, int64, error)
	Update(rps *model.RPS) error
	Delete(id string) error
	UpdateStatus(id, status, reviewedBy string, reviewNotes *string) error
	FindMinimalByID(id string) (*model.RPS, error)
	CountByStatus(status string) (int64, error)
	CountByDosenAndStatus(dosenID, status string) (int64, error)
}

type rpsRepository struct {
	db *gorm.DB
}

func NewRPSRepository(db *gorm.DB) RPSRepository {
	return &rpsRepository{db: db}
}

func (r *rpsRepository) Create(rps *model.RPS) error {
	return r.db.Create(rps).Error
}

func (r *rpsRepository) FindByID(id string) (*model.RPS, error) {
	var rps model.RPS
	err := r.db.Preload("MataKuliah").Preload("Dosen").Preload("Reviewer").
		Preload("CPMK").Preload("RencanaPembelajaran").Preload("BahanBacaan").Preload("Evaluasi").
		Where("id = ?", id).First(&rps).Error
	if err != nil {
		return nil, err
	}
	return &rps, nil
}

func (r *rpsRepository) FindAll(page, limit int, search, mataKuliahID, dosenID, status, tahunAkademik string, semester int, sortBy, sortOrder string) ([]model.RPS, int64, error) {
	var rpsList []model.RPS
	var total int64

	query := r.db.Model(&model.RPS{})

	if search != "" {
		query = query.Where("mata_kuliah_nama LIKE ? OR kode_mk LIKE ? OR dosen_nama LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	if mataKuliahID != "" {
		query = query.Where("mata_kuliah_id = ?", mataKuliahID)
	}
	if dosenID != "" {
		query = query.Where("dosen_id = ?", dosenID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if tahunAkademik != "" {
		query = query.Where("tahun_akademik = ?", tahunAkademik)
	}
	if semester > 0 {
		query = query.Where("semester = ?", semester)
	}

	query.Count(&total)

	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Preload("MataKuliah").Preload("Dosen").
		Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&rpsList).Error
	return rpsList, total, err
}

func (r *rpsRepository) FindByDosenID(dosenID string, page, limit int, status string) ([]model.RPS, int64, error) {
	var rpsList []model.RPS
	var total int64

	query := r.db.Model(&model.RPS{}).Where("dosen_id = ?", dosenID)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("MataKuliah").
		Order("created_at DESC").Offset(offset).Limit(limit).Find(&rpsList).Error
	return rpsList, total, err
}

func (r *rpsRepository) Update(rps *model.RPS) error {
	return r.db.Save(rps).Error
}

func (r *rpsRepository) Delete(id string) error {
	return r.db.Delete(&model.RPS{}, "id = ?", id).Error
}

func (r *rpsRepository) UpdateStatus(id, status, reviewedBy string, reviewNotes *string) error {
	updates := map[string]interface{}{
		"status": status,
	}

	now := time.Now()
	switch status {
	case "submitted":
		updates["submitted_at"] = now
	case "revision":
		updates["reviewed_at"] = now
		updates["reviewer_id"] = reviewedBy
		if reviewNotes != nil {
			updates["review_catatan"] = *reviewNotes
		}
	case "approved":
		updates["reviewed_at"] = now
		updates["approved_at"] = now
		updates["reviewer_id"] = reviewedBy
		if reviewNotes != nil {
			updates["review_catatan"] = *reviewNotes
		}
	case "rejected":
		updates["reviewed_at"] = now
		updates["reviewer_id"] = reviewedBy
		if reviewNotes != nil {
			updates["review_catatan"] = *reviewNotes
		}
	}

	return r.db.Model(&model.RPS{}).Where("id = ?", id).Updates(updates).Error
}

func (r *rpsRepository) FindMinimalByID(id string) (*model.RPS, error) {
	var rps model.RPS
	// Select only the fields we need and avoid preloading associations to prevent unintended writes
	err := r.db.Select("id, mata_kuliah_id, dosen_id").Where("id = ?", id).First(&rps).Error
	if err != nil {
		return nil, err
	}
	// Fetch mata kuliah name separately without using Preload
	var mk model.MataKuliah
	if rps.MataKuliahID != "" {
		if err := r.db.Model(&model.MataKuliah{}).Select("id, nama").Where("id = ?", rps.MataKuliahID).First(&mk).Error; err == nil {
			rps.MataKuliah = mk
		}
	}
	return &rps, nil
}

func (r *rpsRepository) CountByStatus(status string) (int64, error) {
	var count int64
	query := r.db.Model(&model.RPS{})
	if status != "" {
		query = query.Where("status = ?", status)
	} else {
		query = query.Where("status != ?", "draft")
	}
	err := query.Count(&count).Error
	return count, err
}

func (r *rpsRepository) CountByDosenAndStatus(dosenID, status string) (int64, error) {
	var count int64
	query := r.db.Model(&model.RPS{}).Where("dosen_id = ?", dosenID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&count).Error
	return count, err
}
