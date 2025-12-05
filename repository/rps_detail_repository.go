package repository

import (
	"backend-kurikulum-apps/model"

	"gorm.io/gorm"
)

// RPS CPMK Repository
type RPSCPMKRepository interface {
	Create(cpmk *model.RPSCPMK) error
	CreateBatch(cpmks []model.RPSCPMK) error
	FindByRPSID(rpsID string) ([]model.RPSCPMK, error)
	FindByID(id string) (*model.RPSCPMK, error)
	Update(cpmk *model.RPSCPMK) error
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
}

type rpsCPMKRepository struct {
	db *gorm.DB
}

func NewRPSCPMKRepository(db *gorm.DB) RPSCPMKRepository {
	return &rpsCPMKRepository{db: db}
}

func (r *rpsCPMKRepository) Create(cpmk *model.RPSCPMK) error {
	return r.db.Create(cpmk).Error
}

func (r *rpsCPMKRepository) CreateBatch(cpmks []model.RPSCPMK) error {
	return r.db.Create(&cpmks).Error
}

func (r *rpsCPMKRepository) FindByRPSID(rpsID string) ([]model.RPSCPMK, error) {
	var cpmks []model.RPSCPMK
	err := r.db.Where("rps_id = ?", rpsID).Order("urutan ASC").Find(&cpmks).Error
	return cpmks, err
}

func (r *rpsCPMKRepository) FindByID(id string) (*model.RPSCPMK, error) {
	var cpmk model.RPSCPMK
	err := r.db.Where("id = ?", id).First(&cpmk).Error
	if err != nil {
		return nil, err
	}
	return &cpmk, nil
}

func (r *rpsCPMKRepository) Update(cpmk *model.RPSCPMK) error {
	return r.db.Save(cpmk).Error
}

func (r *rpsCPMKRepository) Delete(id string) error {
	return r.db.Delete(&model.RPSCPMK{}, "id = ?", id).Error
}

func (r *rpsCPMKRepository) DeleteByRPSID(rpsID string) error {
	return r.db.Where("rps_id = ?", rpsID).Delete(&model.RPSCPMK{}).Error
}

// RPS Rencana Pembelajaran Repository
type RPSRencanaPembelajaranRepository interface {
	Create(rencana *model.RPSRencanaPembelajaran) error
	CreateBatch(rencanas []model.RPSRencanaPembelajaran) error
	FindByRPSID(rpsID string) ([]model.RPSRencanaPembelajaran, error)
	FindByID(id string) (*model.RPSRencanaPembelajaran, error)
	Update(rencana *model.RPSRencanaPembelajaran) error
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
}

type rpsRencanaPembelajaranRepository struct {
	db *gorm.DB
}

func NewRPSRencanaPembelajaranRepository(db *gorm.DB) RPSRencanaPembelajaranRepository {
	return &rpsRencanaPembelajaranRepository{db: db}
}

func (r *rpsRencanaPembelajaranRepository) Create(rencana *model.RPSRencanaPembelajaran) error {
	return r.db.Create(rencana).Error
}

func (r *rpsRencanaPembelajaranRepository) CreateBatch(rencanas []model.RPSRencanaPembelajaran) error {
	return r.db.Create(&rencanas).Error
}

func (r *rpsRencanaPembelajaranRepository) FindByRPSID(rpsID string) ([]model.RPSRencanaPembelajaran, error) {
	var rencanas []model.RPSRencanaPembelajaran
	err := r.db.Where("rps_id = ?", rpsID).Order("pertemuan ASC").Find(&rencanas).Error
	return rencanas, err
}

func (r *rpsRencanaPembelajaranRepository) FindByID(id string) (*model.RPSRencanaPembelajaran, error) {
	var rencana model.RPSRencanaPembelajaran
	err := r.db.Where("id = ?", id).First(&rencana).Error
	if err != nil {
		return nil, err
	}
	return &rencana, nil
}

func (r *rpsRencanaPembelajaranRepository) Update(rencana *model.RPSRencanaPembelajaran) error {
	return r.db.Save(rencana).Error
}

func (r *rpsRencanaPembelajaranRepository) Delete(id string) error {
	return r.db.Delete(&model.RPSRencanaPembelajaran{}, "id = ?", id).Error
}

func (r *rpsRencanaPembelajaranRepository) DeleteByRPSID(rpsID string) error {
	return r.db.Where("rps_id = ?", rpsID).Delete(&model.RPSRencanaPembelajaran{}).Error
}

// RPS Bahan Bacaan Repository
type RPSBahanBacaanRepository interface {
	Create(bahan *model.RPSBahanBacaan) error
	CreateBatch(bahans []model.RPSBahanBacaan) error
	FindByRPSID(rpsID string) ([]model.RPSBahanBacaan, error)
	FindByID(id string) (*model.RPSBahanBacaan, error)
	Update(bahan *model.RPSBahanBacaan) error
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
}

type rpsBahanBacaanRepository struct {
	db *gorm.DB
}

func NewRPSBahanBacaanRepository(db *gorm.DB) RPSBahanBacaanRepository {
	return &rpsBahanBacaanRepository{db: db}
}

func (r *rpsBahanBacaanRepository) Create(bahan *model.RPSBahanBacaan) error {
	return r.db.Create(bahan).Error
}

func (r *rpsBahanBacaanRepository) CreateBatch(bahans []model.RPSBahanBacaan) error {
	return r.db.Create(&bahans).Error
}

func (r *rpsBahanBacaanRepository) FindByRPSID(rpsID string) ([]model.RPSBahanBacaan, error) {
	var bahans []model.RPSBahanBacaan
	err := r.db.Where("rps_id = ?", rpsID).Order("urutan ASC").Find(&bahans).Error
	return bahans, err
}

func (r *rpsBahanBacaanRepository) FindByID(id string) (*model.RPSBahanBacaan, error) {
	var bahan model.RPSBahanBacaan
	err := r.db.Where("id = ?", id).First(&bahan).Error
	if err != nil {
		return nil, err
	}
	return &bahan, nil
}

func (r *rpsBahanBacaanRepository) Update(bahan *model.RPSBahanBacaan) error {
	return r.db.Save(bahan).Error
}

func (r *rpsBahanBacaanRepository) Delete(id string) error {
	return r.db.Delete(&model.RPSBahanBacaan{}, "id = ?", id).Error
}

func (r *rpsBahanBacaanRepository) DeleteByRPSID(rpsID string) error {
	return r.db.Where("rps_id = ?", rpsID).Delete(&model.RPSBahanBacaan{}).Error
}

// RPS Evaluasi Repository
type RPSEvaluasiRepository interface {
	Create(evaluasi *model.RPSEvaluasi) error
	CreateBatch(evaluasis []model.RPSEvaluasi) error
	FindByRPSID(rpsID string) ([]model.RPSEvaluasi, error)
	FindByID(id string) (*model.RPSEvaluasi, error)
	Update(evaluasi *model.RPSEvaluasi) error
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
	GetTotalBobotByRPSID(rpsID string) (int, error)
}

type rpsEvaluasiRepository struct {
	db *gorm.DB
}

func NewRPSEvaluasiRepository(db *gorm.DB) RPSEvaluasiRepository {
	return &rpsEvaluasiRepository{db: db}
}

func (r *rpsEvaluasiRepository) Create(evaluasi *model.RPSEvaluasi) error {
	return r.db.Create(evaluasi).Error
}

func (r *rpsEvaluasiRepository) CreateBatch(evaluasis []model.RPSEvaluasi) error {
	return r.db.Create(&evaluasis).Error
}

func (r *rpsEvaluasiRepository) FindByRPSID(rpsID string) ([]model.RPSEvaluasi, error) {
	var evaluasis []model.RPSEvaluasi
	err := r.db.Where("rps_id = ?", rpsID).Find(&evaluasis).Error
	return evaluasis, err
}

func (r *rpsEvaluasiRepository) FindByID(id string) (*model.RPSEvaluasi, error) {
	var evaluasi model.RPSEvaluasi
	err := r.db.Where("id = ?", id).First(&evaluasi).Error
	if err != nil {
		return nil, err
	}
	return &evaluasi, nil
}

func (r *rpsEvaluasiRepository) Update(evaluasi *model.RPSEvaluasi) error {
	return r.db.Save(evaluasi).Error
}

func (r *rpsEvaluasiRepository) Delete(id string) error {
	return r.db.Delete(&model.RPSEvaluasi{}, "id = ?", id).Error
}

func (r *rpsEvaluasiRepository) DeleteByRPSID(rpsID string) error {
	return r.db.Where("rps_id = ?", rpsID).Delete(&model.RPSEvaluasi{}).Error
}

func (r *rpsEvaluasiRepository) GetTotalBobotByRPSID(rpsID string) (int, error) {
	var total int
	err := r.db.Model(&model.RPSEvaluasi{}).Where("rps_id = ?", rpsID).Select("COALESCE(SUM(bobot), 0)").Scan(&total).Error
	return total, err
}
