package repository

import (
	"backend-kurikulum-apps/model"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	Create(mk *model.MataKuliah) error
	FindByID(id string) (*model.MataKuliah, error)
	FindByKode(kode string) (*model.MataKuliah, error)
	FindAll(page, limit int, search string, semester int, jenis, status, sortBy, sortOrder string) ([]model.MataKuliah, int64, error)
	FindByDosenID(dosenID string, page, limit int) ([]model.MataKuliah, int64, error)
	Update(mk *model.MataKuliah) error
	ClearDosenPengampu(id string) error
	ClearKoordinator(id string) error
	ClearAllDosen(id string) error
	Delete(id string) error
	SoftDelete(id string) error
	CountBySemester(semester int) (int64, error)
}

type mataKuliahRepository struct {
	db *gorm.DB
}

func NewMataKuliahRepository(db *gorm.DB) MataKuliahRepository {
	return &mataKuliahRepository{db: db}
}

func (r *mataKuliahRepository) Create(mk *model.MataKuliah) error {
	return r.db.Create(mk).Error
}

func (r *mataKuliahRepository) FindByID(id string) (*model.MataKuliah, error) {
	var mk model.MataKuliah
	err := r.db.Preload("DosenPengampu").Preload("Koordinator").Where("id = ?", id).First(&mk).Error
	if err != nil {
		return nil, err
	}
	return &mk, nil
}

func (r *mataKuliahRepository) FindByKode(kode string) (*model.MataKuliah, error) {
	var mk model.MataKuliah
	err := r.db.Where("kode = ?", kode).First(&mk).Error
	if err != nil {
		return nil, err
	}
	return &mk, nil
}

func (r *mataKuliahRepository) FindAll(page, limit int, search string, semester int, jenis, status, sortBy, sortOrder string) ([]model.MataKuliah, int64, error) {
	var mks []model.MataKuliah
	var total int64

	query := r.db.Model(&model.MataKuliah{})

	// Search by kode or nama
	if search != "" {
		query = query.Where("kode LIKE ? OR nama LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Filter by semester
	if semester > 0 {
		query = query.Where("semester = ?", semester)
	}

	// Filter by jenis
	if jenis != "" {
		query = query.Where("jenis = ?", jenis)
	}

	// Filter by status
	if status != "" {
		query = query.Where("status = ?", status)
	} else {
		// Default: exclude deleted
		query = query.Where("status != ?", "dihapus")
	}

	query.Count(&total)

	// Sorting
	if sortBy == "" {
		sortBy = "semester"
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}

	offset := (page - 1) * limit
	err := query.Preload("DosenPengampu").Preload("Koordinator").
		Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&mks).Error
	return mks, total, err
}

func (r *mataKuliahRepository) FindByDosenID(dosenID string, page, limit int) ([]model.MataKuliah, int64, error) {
	var mks []model.MataKuliah
	var total int64

	query := r.db.Model(&model.MataKuliah{}).
		Where("dosen_pengampu_id = ? OR koordinator_id = ?", dosenID, dosenID).
		Where("status = ?", "aktif")

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("DosenPengampu").Preload("Koordinator").
		Order("semester asc, kode asc").Offset(offset).Limit(limit).Find(&mks).Error
	return mks, total, err
}

func (r *mataKuliahRepository) Update(mk *model.MataKuliah) error {
	return r.db.Save(mk).Error
}

func (r *mataKuliahRepository) ClearDosenPengampu(id string) error {
	return r.db.Model(&model.MataKuliah{}).Where("id = ?", id).Update("dosen_pengampu_id", nil).Error
}

func (r *mataKuliahRepository) ClearKoordinator(id string) error {
	return r.db.Model(&model.MataKuliah{}).Where("id = ?", id).Update("koordinator_id", nil).Error
}

func (r *mataKuliahRepository) ClearAllDosen(id string) error {
	return r.db.Model(&model.MataKuliah{}).Where("id = ?", id).Updates(map[string]interface{}{
		"dosen_pengampu_id": nil,
		"koordinator_id":    nil,
	}).Error
}

func (r *mataKuliahRepository) Delete(id string) error {
	return r.db.Delete(&model.MataKuliah{}, "id = ?", id).Error
}

func (r *mataKuliahRepository) SoftDelete(id string) error {
	return r.db.Model(&model.MataKuliah{}).Where("id = ?", id).Update("status", "dihapus").Error
}

func (r *mataKuliahRepository) CountBySemester(semester int) (int64, error) {
	var count int64
	query := r.db.Model(&model.MataKuliah{}).Where("status = ?", "aktif")
	if semester > 0 {
		query = query.Where("semester = ?", semester)
	}
	err := query.Count(&count).Error
	return count, err
}
