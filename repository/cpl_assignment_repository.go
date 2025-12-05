package repository

import (
	"backend-kurikulum-apps/model"
	"time"

	"gorm.io/gorm"
)

type CPLAssignmentRepository interface {
	Create(assignment *model.CPLAssignment) error
	FindByID(id string) (*model.CPLAssignment, error)
	FindAll(page, limit int, cplID, dosenID, status, sortBy, sortOrder string) ([]model.CPLAssignment, int64, error)
	FindByDosenID(dosenID string, page, limit int, status string) ([]model.CPLAssignment, int64, error)
	Update(assignment *model.CPLAssignment) error
	Delete(id string) error
	UpdateStatus(id, status string, rejectionReason *string) error
	CountByStatus(status string) (int64, error)
	CountByDosenAndStatus(dosenID, status string) (int64, error)
	CheckDuplicateAssignment(cplID, dosenID, mataKuliahID string) (bool, error)
}

type cplAssignmentRepository struct {
	db *gorm.DB
}

func NewCPLAssignmentRepository(db *gorm.DB) CPLAssignmentRepository {
	return &cplAssignmentRepository{db: db}
}

func (r *cplAssignmentRepository) Create(assignment *model.CPLAssignment) error {
	return r.db.Create(assignment).Error
}

func (r *cplAssignmentRepository) FindByID(id string) (*model.CPLAssignment, error) {
	var assignment model.CPLAssignment
	err := r.db.Preload("CPL").Preload("Dosen").Preload("MataKuliahRef").Preload("Assigner").
		Where("id = ?", id).First(&assignment).Error
	if err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (r *cplAssignmentRepository) FindAll(page, limit int, cplID, dosenID, status, sortBy, sortOrder string) ([]model.CPLAssignment, int64, error) {
	var assignments []model.CPLAssignment
	var total int64

	query := r.db.Model(&model.CPLAssignment{})

	if cplID != "" {
		query = query.Where("cpl_id = ?", cplID)
	}
	if dosenID != "" {
		query = query.Where("dosen_id = ?", dosenID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if sortBy == "" {
		sortBy = "assigned_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Preload("CPL").Preload("Dosen").Preload("MataKuliahRef").
		Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&assignments).Error
	return assignments, total, err
}

func (r *cplAssignmentRepository) FindByDosenID(dosenID string, page, limit int, status string) ([]model.CPLAssignment, int64, error) {
	var assignments []model.CPLAssignment
	var total int64

	query := r.db.Model(&model.CPLAssignment{}).Where("dosen_id = ?", dosenID)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("CPL").Preload("MataKuliahRef").
		Order("assigned_at DESC").Offset(offset).Limit(limit).Find(&assignments).Error
	return assignments, total, err
}

func (r *cplAssignmentRepository) Update(assignment *model.CPLAssignment) error {
	return r.db.Save(assignment).Error
}

func (r *cplAssignmentRepository) Delete(id string) error {
	return r.db.Delete(&model.CPLAssignment{}, "id = ?", id).Error
}

func (r *cplAssignmentRepository) UpdateStatus(id, status string, rejectionReason *string) error {
	updates := map[string]interface{}{
		"status": status,
	}

	now := time.Now()
	switch status {
	case "accepted":
		updates["accepted_at"] = now
	case "done":
		updates["completed_at"] = now
	case "rejected":
		if rejectionReason != nil {
			updates["rejection_reason"] = *rejectionReason
		}
	}

	return r.db.Model(&model.CPLAssignment{}).Where("id = ?", id).Updates(updates).Error
}

func (r *cplAssignmentRepository) CountByStatus(status string) (int64, error) {
	var count int64
	query := r.db.Model(&model.CPLAssignment{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&count).Error
	return count, err
}

func (r *cplAssignmentRepository) CountByDosenAndStatus(dosenID, status string) (int64, error) {
	var count int64
	query := r.db.Model(&model.CPLAssignment{}).Where("dosen_id = ?", dosenID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&count).Error
	return count, err
}

func (r *cplAssignmentRepository) CheckDuplicateAssignment(cplID, dosenID, mataKuliahID string) (bool, error) {
	var count int64
	query := r.db.Model(&model.CPLAssignment{}).
		Where("cpl_id = ? AND dosen_id = ? AND mata_kuliah_id = ?", cplID, dosenID, mataKuliahID).
		Where("status NOT IN ?", []string{"rejected", "cancelled"})
	err := query.Count(&count).Error
	return count > 0, err
}
