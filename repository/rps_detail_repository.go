package repository

import (
	"backend-kurikulum-apps/model"

	"gorm.io/gorm"
)

// RPS CPMK Repository
type RPSCPMKRepository interface {
	Create(cpmk *model.RPSCPMK) error
	CreateBatch(cpmks []model.RPSCPMK) error
	FindAll() ([]model.RPSCPMK, error)
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

func (r *rpsCPMKRepository) FindAll() ([]model.RPSCPMK, error) {
	var cpmks []model.RPSCPMK
	err := r.db.Order("urutan ASC").Find(&cpmks).Error
	return cpmks, err
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
	err := r.db.Where("rps_id = ?", rpsID).Order("minggu_ke ASC").Find(&rencanas).Error
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


// ============ Sub-CPMK Repository ============

type SubCPMKRepository interface {
	Create(subCpmk *model.SubCPMK) error
	CreateBatch(subCpmks []model.SubCPMK) error
	FindByCPMKID(cpmkID string) ([]model.SubCPMK, error)
	FindByID(id string) (*model.SubCPMK, error)
	FindByIDs(ids []string) ([]model.SubCPMK, error)
	Update(subCpmk *model.SubCPMK) error
	Delete(id string) error
	DeleteByCPMKID(cpmkID string) error
}

type subCPMKRepository struct {
	db *gorm.DB
}

func NewSubCPMKRepository(db *gorm.DB) SubCPMKRepository {
	return &subCPMKRepository{db: db}
}

func (r *subCPMKRepository) Create(subCpmk *model.SubCPMK) error {
	return r.db.Create(subCpmk).Error
}

func (r *subCPMKRepository) CreateBatch(subCpmks []model.SubCPMK) error {
	return r.db.Create(&subCpmks).Error
}

func (r *subCPMKRepository) FindByCPMKID(cpmkID string) ([]model.SubCPMK, error) {
	var subCpmks []model.SubCPMK
	err := r.db.Where("cpmk_id = ?", cpmkID).Order("urutan ASC").Find(&subCpmks).Error
	return subCpmks, err
}

func (r *subCPMKRepository) FindByID(id string) (*model.SubCPMK, error) {
	var subCpmk model.SubCPMK
	err := r.db.Where("id = ?", id).First(&subCpmk).Error
	if err != nil {
		return nil, err
	}
	return &subCpmk, nil
}

func (r *subCPMKRepository) FindByIDs(ids []string) ([]model.SubCPMK, error) {
	var subCpmks []model.SubCPMK
	err := r.db.Where("id IN ?", ids).Find(&subCpmks).Error
	return subCpmks, err
}

func (r *subCPMKRepository) Update(subCpmk *model.SubCPMK) error {
	return r.db.Save(subCpmk).Error
}

func (r *subCPMKRepository) Delete(id string) error {
	return r.db.Delete(&model.SubCPMK{}, "id = ?", id).Error
}

func (r *subCPMKRepository) DeleteByCPMKID(cpmkID string) error {
	return r.db.Where("cpmk_id = ?", cpmkID).Delete(&model.SubCPMK{}).Error
}

// ============ RPS Rencana Tugas Repository ============

type RPSRencanaTugasRepository interface {
	Create(tugas *model.RPSRencanaTugas) error
	CreateBatch(tugasList []model.RPSRencanaTugas) error
	FindByRPSID(rpsID string) ([]model.RPSRencanaTugas, error)
	FindByID(id string) (*model.RPSRencanaTugas, error)
	Update(tugas *model.RPSRencanaTugas) error
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
	GetTotalBobotByRPSID(rpsID string) (int, error)
}

type rpsRencanaTugasRepository struct {
	db *gorm.DB
}

func NewRPSRencanaTugasRepository(db *gorm.DB) RPSRencanaTugasRepository {
	return &rpsRencanaTugasRepository{db: db}
}

func (r *rpsRencanaTugasRepository) Create(tugas *model.RPSRencanaTugas) error {
	return r.db.Create(tugas).Error
}

func (r *rpsRencanaTugasRepository) CreateBatch(tugasList []model.RPSRencanaTugas) error {
	return r.db.Create(&tugasList).Error
}

func (r *rpsRencanaTugasRepository) FindByRPSID(rpsID string) ([]model.RPSRencanaTugas, error) {
	var tugasList []model.RPSRencanaTugas
	err := r.db.Where("rps_id = ?", rpsID).Order("nomor_tugas ASC").Find(&tugasList).Error
	return tugasList, err
}

func (r *rpsRencanaTugasRepository) FindByID(id string) (*model.RPSRencanaTugas, error) {
	var tugas model.RPSRencanaTugas
	err := r.db.Where("id = ?", id).First(&tugas).Error
	if err != nil {
		return nil, err
	}
	return &tugas, nil
}

func (r *rpsRencanaTugasRepository) Update(tugas *model.RPSRencanaTugas) error {
	return r.db.Save(tugas).Error
}

func (r *rpsRencanaTugasRepository) Delete(id string) error {
	return r.db.Delete(&model.RPSRencanaTugas{}, "id = ?", id).Error
}

func (r *rpsRencanaTugasRepository) DeleteByRPSID(rpsID string) error {
	return r.db.Where("rps_id = ?", rpsID).Delete(&model.RPSRencanaTugas{}).Error
}

func (r *rpsRencanaTugasRepository) GetTotalBobotByRPSID(rpsID string) (int, error) {
	var total int
	err := r.db.Model(&model.RPSRencanaTugas{}).Where("rps_id = ?", rpsID).Select("COALESCE(SUM(bobot), 0)").Scan(&total).Error
	return total, err
}

// ============ RPS Analisis Ketercapaian CPL Repository ============

type RPSAnalisisKetercapaianCPLRepository interface {
	Create(analisis *model.RPSAnalisisKetercapaianCPL) error
	CreateBatch(analisisList []model.RPSAnalisisKetercapaianCPL) error
	FindByRPSID(rpsID string) ([]model.RPSAnalisisKetercapaianCPL, error)
	FindByID(id string) (*model.RPSAnalisisKetercapaianCPL, error)
	FindByCPLID(cplID string) ([]model.RPSAnalisisKetercapaianCPL, error)
	Update(analisis *model.RPSAnalisisKetercapaianCPL) error
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
}

type rpsAnalisisKetercapaianCPLRepository struct {
	db *gorm.DB
}

func NewRPSAnalisisKetercapaianCPLRepository(db *gorm.DB) RPSAnalisisKetercapaianCPLRepository {
	return &rpsAnalisisKetercapaianCPLRepository{db: db}
}

func (r *rpsAnalisisKetercapaianCPLRepository) Create(analisis *model.RPSAnalisisKetercapaianCPL) error {
	return r.db.Create(analisis).Error
}

func (r *rpsAnalisisKetercapaianCPLRepository) CreateBatch(analisisList []model.RPSAnalisisKetercapaianCPL) error {
	return r.db.Create(&analisisList).Error
}

func (r *rpsAnalisisKetercapaianCPLRepository) FindByRPSID(rpsID string) ([]model.RPSAnalisisKetercapaianCPL, error) {
	var analisisList []model.RPSAnalisisKetercapaianCPL
	err := r.db.Preload("CPL").Where("rps_id = ?", rpsID).Order("minggu_mulai ASC").Find(&analisisList).Error
	return analisisList, err
}

func (r *rpsAnalisisKetercapaianCPLRepository) FindByID(id string) (*model.RPSAnalisisKetercapaianCPL, error) {
	var analisis model.RPSAnalisisKetercapaianCPL
	err := r.db.Preload("CPL").Where("id = ?", id).First(&analisis).Error
	if err != nil {
		return nil, err
	}
	return &analisis, nil
}

func (r *rpsAnalisisKetercapaianCPLRepository) FindByCPLID(cplID string) ([]model.RPSAnalisisKetercapaianCPL, error) {
	var analisisList []model.RPSAnalisisKetercapaianCPL
	err := r.db.Where("cpl_id = ?", cplID).Find(&analisisList).Error
	return analisisList, err
}

func (r *rpsAnalisisKetercapaianCPLRepository) Update(analisis *model.RPSAnalisisKetercapaianCPL) error {
	return r.db.Save(analisis).Error
}

func (r *rpsAnalisisKetercapaianCPLRepository) Delete(id string) error {
	return r.db.Delete(&model.RPSAnalisisKetercapaianCPL{}, "id = ?", id).Error
}

func (r *rpsAnalisisKetercapaianCPLRepository) DeleteByRPSID(rpsID string) error {
	return r.db.Where("rps_id = ?", rpsID).Delete(&model.RPSAnalisisKetercapaianCPL{}).Error
}


// ============ CPMK CPL Mapping Repository ============

type CPMKCPLMappingRepository interface {
	Create(mapping *model.CPMKCPLMapping) error
	CreateBatch(mappings []model.CPMKCPLMapping) error
	FindByCPMKID(cpmkID string) ([]model.CPMKCPLMapping, error)
	FindByCPLID(cplID string) ([]model.CPMKCPLMapping, error)
	Delete(id string) error
	DeleteByCPMKID(cpmkID string) error
	SyncCPLsForCPMK(cpmkID string, cplIDs []string) error
}

type cpmkCPLMappingRepository struct {
	db *gorm.DB
}

func NewCPMKCPLMappingRepository(db *gorm.DB) CPMKCPLMappingRepository {
	return &cpmkCPLMappingRepository{db: db}
}

func (r *cpmkCPLMappingRepository) Create(mapping *model.CPMKCPLMapping) error {
	return r.db.Create(mapping).Error
}

func (r *cpmkCPLMappingRepository) CreateBatch(mappings []model.CPMKCPLMapping) error {
	return r.db.Create(&mappings).Error
}

func (r *cpmkCPLMappingRepository) FindByCPMKID(cpmkID string) ([]model.CPMKCPLMapping, error) {
	var mappings []model.CPMKCPLMapping
	err := r.db.Preload("CPL").Where("cpmk_id = ?", cpmkID).Find(&mappings).Error
	return mappings, err
}

func (r *cpmkCPLMappingRepository) FindByCPLID(cplID string) ([]model.CPMKCPLMapping, error) {
	var mappings []model.CPMKCPLMapping
	err := r.db.Preload("CPMK").Where("cpl_id = ?", cplID).Find(&mappings).Error
	return mappings, err
}

func (r *cpmkCPLMappingRepository) Delete(id string) error {
	return r.db.Delete(&model.CPMKCPLMapping{}, "id = ?", id).Error
}

func (r *cpmkCPLMappingRepository) DeleteByCPMKID(cpmkID string) error {
	return r.db.Where("cpmk_id = ?", cpmkID).Delete(&model.CPMKCPLMapping{}).Error
}

func (r *cpmkCPLMappingRepository) SyncCPLsForCPMK(cpmkID string, cplIDs []string) error {
	// Delete existing mappings
	if err := r.DeleteByCPMKID(cpmkID); err != nil {
		return err
	}

	// Create new mappings
	if len(cplIDs) > 0 {
		var mappings []model.CPMKCPLMapping
		for _, cplID := range cplIDs {
			mappings = append(mappings, model.CPMKCPLMapping{
				CPMKID: cpmkID,
				CPLID:  cplID,
			})
		}
		return r.CreateBatch(mappings)
	}
	return nil
}
