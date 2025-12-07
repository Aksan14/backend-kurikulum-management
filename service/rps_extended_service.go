package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"encoding/json"
	"errors"
	"strings"
)

// ============ Sub-CPMK Service ============

type SubCPMKService interface {
	GetByCPMKID(cpmkID string) ([]dto.SubCPMKResponse, error)
	GetByID(id string) (*dto.SubCPMKResponse, error)
	Create(cpmkID string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error)
	Update(id string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error)
	Delete(id string) error
}

type subCPMKService struct {
	subCpmkRepo repository.SubCPMKRepository
	cpmkRepo    repository.RPSCPMKRepository
}

func NewSubCPMKService(
	subCpmkRepo repository.SubCPMKRepository,
	cpmkRepo repository.RPSCPMKRepository,
) SubCPMKService {
	return &subCPMKService{
		subCpmkRepo: subCpmkRepo,
		cpmkRepo:    cpmkRepo,
	}
}

func (s *subCPMKService) GetByCPMKID(cpmkID string) ([]dto.SubCPMKResponse, error) {
	subCpmks, err := s.subCpmkRepo.FindByCPMKID(cpmkID)
	if err != nil {
		return nil, err
	}

	var responses []dto.SubCPMKResponse
	for _, subCpmk := range subCpmks {
		responses = append(responses, toSubCPMKResponse(&subCpmk))
	}
	return responses, nil
}

func (s *subCPMKService) GetByID(id string) (*dto.SubCPMKResponse, error) {
	subCpmk, err := s.subCpmkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Sub-CPMK tidak ditemukan")
	}
	resp := toSubCPMKResponse(subCpmk)
	return &resp, nil
}

func (s *subCPMKService) Create(cpmkID string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error) {
	// Verify CPMK exists
	_, err := s.cpmkRepo.FindByID(cpmkID)
	if err != nil {
		return nil, errors.New("CPMK tidak ditemukan")
	}

	subCpmk := &model.SubCPMK{
		CPMKID:    cpmkID,
		Kode:      req.Kode,
		Deskripsi: req.Deskripsi,
		Urutan:    req.Urutan,
	}

	if err := s.subCpmkRepo.Create(subCpmk); err != nil {
		return nil, errors.New("gagal membuat Sub-CPMK")
	}

	resp := toSubCPMKResponse(subCpmk)
	return &resp, nil
}

func (s *subCPMKService) Update(id string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error) {
	subCpmk, err := s.subCpmkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Sub-CPMK tidak ditemukan")
	}

	subCpmk.Kode = req.Kode
	subCpmk.Deskripsi = req.Deskripsi
	subCpmk.Urutan = req.Urutan

	if err := s.subCpmkRepo.Update(subCpmk); err != nil {
		return nil, errors.New("gagal mengupdate Sub-CPMK")
	}

	resp := toSubCPMKResponse(subCpmk)
	return &resp, nil
}

func (s *subCPMKService) Delete(id string) error {
	return s.subCpmkRepo.Delete(id)
}

func toSubCPMKResponse(subCpmk *model.SubCPMK) dto.SubCPMKResponse {
	return dto.SubCPMKResponse{
		ID:        subCpmk.ID,
		CPMKID:    subCpmk.CPMKID,
		Kode:      subCpmk.Kode,
		Deskripsi: subCpmk.Deskripsi,
		Urutan:    subCpmk.Urutan,
		CreatedAt: subCpmk.CreatedAt,
		UpdatedAt: subCpmk.UpdatedAt,
	}
}

// ============ RPS Rencana Tugas Service ============

type RPSRencanaTugasService interface {
	GetByRPSID(rpsID string) ([]dto.RPSRencanaTugasResponse, error)
	GetByID(id string) (*dto.RPSRencanaTugasResponse, error)
	Create(rpsID string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error)
	Update(id string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error)
	Delete(id string) error
}

type rpsRencanaTugasService struct {
	tugasRepo repository.RPSRencanaTugasRepository
	rpsRepo   repository.RPSRepository
}

func NewRPSRencanaTugasService(
	tugasRepo repository.RPSRencanaTugasRepository,
	rpsRepo repository.RPSRepository,
) RPSRencanaTugasService {
	return &rpsRencanaTugasService{
		tugasRepo: tugasRepo,
		rpsRepo:   rpsRepo,
	}
}

func (s *rpsRencanaTugasService) GetByRPSID(rpsID string) ([]dto.RPSRencanaTugasResponse, error) {
	tugasList, err := s.tugasRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSRencanaTugasResponse
	for _, tugas := range tugasList {
		responses = append(responses, toRPSRencanaTugasResponse(&tugas))
	}
	return responses, nil
}

func (s *rpsRencanaTugasService) GetByID(id string) (*dto.RPSRencanaTugasResponse, error) {
	tugas, err := s.tugasRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("rencana tugas tidak ditemukan")
	}
	resp := toRPSRencanaTugasResponse(tugas)
	return &resp, nil
}

func (s *rpsRencanaTugasService) Create(rpsID string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error) {
	// Verify RPS exists
	_, err := s.rpsRepo.FindByID(rpsID)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	// Check total bobot
	currentTotal, _ := s.tugasRepo.GetTotalBobotByRPSID(rpsID)
	if currentTotal+req.Bobot > 100 {
		return nil, errors.New("total bobot tugas tidak boleh lebih dari 100%")
	}

	jenisTugas := strings.ToLower(req.JenisTugas)
	if jenisTugas == "" {
		jenisTugas = "individu"
	}

	tugas := &model.RPSRencanaTugas{
		RPSID:                 rpsID,
		NomorTugas:            req.NomorTugas,
		Judul:                 req.Judul,
		IndikatorKeberhasilan: req.IndikatorKeberhasilan,
		BatasWaktuMinggu:      req.BatasWaktuMinggu,
		PetunjukPengerjaan:    req.PetunjukPengerjaan,
		JenisTugas:            jenisTugas,
		LuaranTugas:           req.LuaranTugas,
		KriteriaPenilaian:     req.KriteriaPenilaian,
		TeknikPenilaian:       req.TeknikPenilaian,
		Bobot:                 req.Bobot,
	}

	if err := s.tugasRepo.Create(tugas); err != nil {
		return nil, errors.New("gagal membuat rencana tugas")
	}

	resp := toRPSRencanaTugasResponse(tugas)
	return &resp, nil
}

func (s *rpsRencanaTugasService) Update(id string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error) {
	tugas, err := s.tugasRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("rencana tugas tidak ditemukan")
	}

	// Check total bobot (excluding current)
	currentTotal, _ := s.tugasRepo.GetTotalBobotByRPSID(tugas.RPSID)
	if (currentTotal-tugas.Bobot)+req.Bobot > 100 {
		return nil, errors.New("total bobot tugas tidak boleh lebih dari 100%")
	}

	tugas.NomorTugas = req.NomorTugas
	tugas.Judul = req.Judul
	tugas.IndikatorKeberhasilan = req.IndikatorKeberhasilan
	tugas.BatasWaktuMinggu = req.BatasWaktuMinggu
	tugas.PetunjukPengerjaan = req.PetunjukPengerjaan
	if req.JenisTugas != "" {
		tugas.JenisTugas = strings.ToLower(req.JenisTugas)
	}
	tugas.LuaranTugas = req.LuaranTugas
	tugas.KriteriaPenilaian = req.KriteriaPenilaian
	tugas.TeknikPenilaian = req.TeknikPenilaian
	tugas.Bobot = req.Bobot

	if err := s.tugasRepo.Update(tugas); err != nil {
		return nil, errors.New("gagal mengupdate rencana tugas")
	}

	resp := toRPSRencanaTugasResponse(tugas)
	return &resp, nil
}

func (s *rpsRencanaTugasService) Delete(id string) error {
	return s.tugasRepo.Delete(id)
}

func toRPSRencanaTugasResponse(tugas *model.RPSRencanaTugas) dto.RPSRencanaTugasResponse {
	return dto.RPSRencanaTugasResponse{
		ID:                    tugas.ID,
		RPSID:                 tugas.RPSID,
		NomorTugas:            tugas.NomorTugas,
		Judul:                 tugas.Judul,
		IndikatorKeberhasilan: tugas.IndikatorKeberhasilan,
		BatasWaktuMinggu:      tugas.BatasWaktuMinggu,
		PetunjukPengerjaan:    tugas.PetunjukPengerjaan,
		JenisTugas:            tugas.JenisTugas,
		LuaranTugas:           tugas.LuaranTugas,
		KriteriaPenilaian:     tugas.KriteriaPenilaian,
		TeknikPenilaian:       tugas.TeknikPenilaian,
		Bobot:                 tugas.Bobot,
		CreatedAt:             tugas.CreatedAt,
		UpdatedAt:             tugas.UpdatedAt,
	}
}

// ============ RPS Analisis Ketercapaian CPL Service ============

type RPSAnalisisKetercapaianCPLService interface {
	GetByRPSID(rpsID string) ([]dto.RPSAnalisisKetercapaianCPLResponse, error)
	GetByID(id string) (*dto.RPSAnalisisKetercapaianCPLResponse, error)
	Create(rpsID string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error)
	Update(id string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error)
	Delete(id string) error
}

type rpsAnalisisKetercapaianCPLService struct {
	analisisRepo repository.RPSAnalisisKetercapaianCPLRepository
	rpsRepo      repository.RPSRepository
}

func NewRPSAnalisisKetercapaianCPLService(
	analisisRepo repository.RPSAnalisisKetercapaianCPLRepository,
	rpsRepo repository.RPSRepository,
) RPSAnalisisKetercapaianCPLService {
	return &rpsAnalisisKetercapaianCPLService{
		analisisRepo: analisisRepo,
		rpsRepo:      rpsRepo,
	}
}

func (s *rpsAnalisisKetercapaianCPLService) GetByRPSID(rpsID string) ([]dto.RPSAnalisisKetercapaianCPLResponse, error) {
	analisisList, err := s.analisisRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSAnalisisKetercapaianCPLResponse
	for _, analisis := range analisisList {
		responses = append(responses, toRPSAnalisisKetercapaianCPLResponse(&analisis))
	}
	return responses, nil
}

func (s *rpsAnalisisKetercapaianCPLService) GetByID(id string) (*dto.RPSAnalisisKetercapaianCPLResponse, error) {
	analisis, err := s.analisisRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("analisis ketercapaian CPL tidak ditemukan")
	}
	resp := toRPSAnalisisKetercapaianCPLResponse(analisis)
	return &resp, nil
}

func (s *rpsAnalisisKetercapaianCPLService) Create(rpsID string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error) {
	// Verify RPS exists
	_, err := s.rpsRepo.FindByID(rpsID)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	cpmkIDsJSON, _ := json.Marshal(req.CPMKIDs)
	subCpmkIDsJSON, _ := json.Marshal(req.SubCPMKIDs)

	analisis := &model.RPSAnalisisKetercapaianCPL{
		RPSID:           rpsID,
		MingguMulai:     req.MingguMulai,
		MingguSelesai:   req.MingguSelesai,
		CPLID:           req.CPLID,
		CPMKIDs:         cpmkIDsJSON,
		SubCPMKIDs:      subCpmkIDsJSON,
		TopikMateri:     req.TopikMateri,
		JenisAssessment: req.JenisAssessment,
		BobotKontribusi: req.BobotKontribusi,
	}

	if err := s.analisisRepo.Create(analisis); err != nil {
		return nil, errors.New("gagal membuat analisis ketercapaian CPL")
	}

	// Reload to get CPL relation
	analisis, _ = s.analisisRepo.FindByID(analisis.ID)
	resp := toRPSAnalisisKetercapaianCPLResponse(analisis)
	return &resp, nil
}

func (s *rpsAnalisisKetercapaianCPLService) Update(id string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error) {
	analisis, err := s.analisisRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("analisis ketercapaian CPL tidak ditemukan")
	}

	cpmkIDsJSON, _ := json.Marshal(req.CPMKIDs)
	subCpmkIDsJSON, _ := json.Marshal(req.SubCPMKIDs)

	analisis.MingguMulai = req.MingguMulai
	analisis.MingguSelesai = req.MingguSelesai
	analisis.CPLID = req.CPLID
	analisis.CPMKIDs = cpmkIDsJSON
	analisis.SubCPMKIDs = subCpmkIDsJSON
	analisis.TopikMateri = req.TopikMateri
	analisis.JenisAssessment = req.JenisAssessment
	analisis.BobotKontribusi = req.BobotKontribusi

	if err := s.analisisRepo.Update(analisis); err != nil {
		return nil, errors.New("gagal mengupdate analisis ketercapaian CPL")
	}

	// Reload to get CPL relation
	analisis, _ = s.analisisRepo.FindByID(analisis.ID)
	resp := toRPSAnalisisKetercapaianCPLResponse(analisis)
	return &resp, nil
}

func (s *rpsAnalisisKetercapaianCPLService) Delete(id string) error {
	return s.analisisRepo.Delete(id)
}

func toRPSAnalisisKetercapaianCPLResponse(analisis *model.RPSAnalisisKetercapaianCPL) dto.RPSAnalisisKetercapaianCPLResponse {
	var cpmkIDs []string
	var subCpmkIDs []string
	json.Unmarshal(analisis.CPMKIDs, &cpmkIDs)
	json.Unmarshal(analisis.SubCPMKIDs, &subCpmkIDs)

	resp := dto.RPSAnalisisKetercapaianCPLResponse{
		ID:              analisis.ID,
		RPSID:           analisis.RPSID,
		MingguMulai:     analisis.MingguMulai,
		MingguSelesai:   analisis.MingguSelesai,
		CPLID:           analisis.CPLID,
		CPMKIDs:         cpmkIDs,
		SubCPMKIDs:      subCpmkIDs,
		TopikMateri:     analisis.TopikMateri,
		JenisAssessment: analisis.JenisAssessment,
		BobotKontribusi: analisis.BobotKontribusi,
		CreatedAt:       analisis.CreatedAt,
		UpdatedAt:       analisis.UpdatedAt,
	}

	if analisis.CPL.ID != "" {
		resp.CPL = &dto.CPLSimpleResponse{
			ID:   analisis.CPL.ID,
			Kode: analisis.CPL.Kode,
			Nama: analisis.CPL.Nama,
		}
	}

	return resp
}

// ============ RPS Skala Penilaian Service ============

type RPSSkalaPenilaianService interface {
	GetByRPSID(rpsID string) ([]dto.RPSSkalaPenilaianResponse, error)
	GetByID(id string) (*dto.RPSSkalaPenilaianResponse, error)
	Create(rpsID string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error)
	CreateBatch(rpsID string, req dto.RPSBatchSkalaPenilaianRequest) ([]dto.RPSSkalaPenilaianResponse, error)
	CreateDefault(rpsID string) ([]dto.RPSSkalaPenilaianResponse, error)
	Update(id string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error)
	Delete(id string) error
	DeleteByRPSID(rpsID string) error
}

type rpsSkalaPenilaianService struct {
	skalaRepo repository.RPSSkalaPenilaianRepository
	rpsRepo   repository.RPSRepository
}

func NewRPSSkalaPenilaianService(
	skalaRepo repository.RPSSkalaPenilaianRepository,
	rpsRepo repository.RPSRepository,
) RPSSkalaPenilaianService {
	return &rpsSkalaPenilaianService{
		skalaRepo: skalaRepo,
		rpsRepo:   rpsRepo,
	}
}

func (s *rpsSkalaPenilaianService) GetByRPSID(rpsID string) ([]dto.RPSSkalaPenilaianResponse, error) {
	skalaList, err := s.skalaRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSSkalaPenilaianResponse
	for _, skala := range skalaList {
		responses = append(responses, toRPSSkalaPenilaianResponse(&skala))
	}
	return responses, nil
}

func (s *rpsSkalaPenilaianService) GetByID(id string) (*dto.RPSSkalaPenilaianResponse, error) {
	skala, err := s.skalaRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("skala penilaian tidak ditemukan")
	}
	resp := toRPSSkalaPenilaianResponse(skala)
	return &resp, nil
}

func (s *rpsSkalaPenilaianService) Create(rpsID string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error) {
	// Verify RPS exists
	_, err := s.rpsRepo.FindByID(rpsID)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	skala := &model.RPSSkalaPenilaian{
		RPSID:      rpsID,
		NilaiMin:   req.NilaiMin,
		NilaiMax:   req.NilaiMax,
		HurufMutu:  req.HurufMutu,
		BobotNilai: req.BobotNilai,
		IsLulus:    req.IsLulus,
	}

	if err := s.skalaRepo.Create(skala); err != nil {
		return nil, errors.New("gagal membuat skala penilaian")
	}

	resp := toRPSSkalaPenilaianResponse(skala)
	return &resp, nil
}

func (s *rpsSkalaPenilaianService) CreateBatch(rpsID string, req dto.RPSBatchSkalaPenilaianRequest) ([]dto.RPSSkalaPenilaianResponse, error) {
	// Verify RPS exists
	_, err := s.rpsRepo.FindByID(rpsID)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	// Delete existing skala
	s.skalaRepo.DeleteByRPSID(rpsID)

	var skalaList []model.RPSSkalaPenilaian
	for _, skalaReq := range req.SkalaPenilaian {
		skalaList = append(skalaList, model.RPSSkalaPenilaian{
			RPSID:      rpsID,
			NilaiMin:   skalaReq.NilaiMin,
			NilaiMax:   skalaReq.NilaiMax,
			HurufMutu:  skalaReq.HurufMutu,
			BobotNilai: skalaReq.BobotNilai,
			IsLulus:    skalaReq.IsLulus,
		})
	}

	if err := s.skalaRepo.CreateBatch(skalaList); err != nil {
		return nil, errors.New("gagal membuat skala penilaian")
	}

	// Return created items
	return s.GetByRPSID(rpsID)
}

func (s *rpsSkalaPenilaianService) CreateDefault(rpsID string) ([]dto.RPSSkalaPenilaianResponse, error) {
	// Verify RPS exists
	_, err := s.rpsRepo.FindByID(rpsID)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	// Delete existing skala
	s.skalaRepo.DeleteByRPSID(rpsID)

	// Create default
	if err := s.skalaRepo.CreateDefaultSkala(rpsID); err != nil {
		return nil, errors.New("gagal membuat skala penilaian default")
	}

	return s.GetByRPSID(rpsID)
}

func (s *rpsSkalaPenilaianService) Update(id string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error) {
	skala, err := s.skalaRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("skala penilaian tidak ditemukan")
	}

	skala.NilaiMin = req.NilaiMin
	skala.NilaiMax = req.NilaiMax
	skala.HurufMutu = req.HurufMutu
	skala.BobotNilai = req.BobotNilai
	skala.IsLulus = req.IsLulus

	if err := s.skalaRepo.Update(skala); err != nil {
		return nil, errors.New("gagal mengupdate skala penilaian")
	}

	resp := toRPSSkalaPenilaianResponse(skala)
	return &resp, nil
}

func (s *rpsSkalaPenilaianService) Delete(id string) error {
	return s.skalaRepo.Delete(id)
}

func (s *rpsSkalaPenilaianService) DeleteByRPSID(rpsID string) error {
	return s.skalaRepo.DeleteByRPSID(rpsID)
}

func toRPSSkalaPenilaianResponse(skala *model.RPSSkalaPenilaian) dto.RPSSkalaPenilaianResponse {
	return dto.RPSSkalaPenilaianResponse{
		ID:         skala.ID,
		RPSID:      skala.RPSID,
		NilaiMin:   skala.NilaiMin,
		NilaiMax:   skala.NilaiMax,
		HurufMutu:  skala.HurufMutu,
		BobotNilai: skala.BobotNilai,
		IsLulus:    skala.IsLulus,
		CreatedAt:  skala.CreatedAt,
	}
}

// ============ RPSExtendedService - Combined Interface ============

type RPSExtendedService interface {
	// Sub-CPMK Methods
	AddSubCPMK(cpmkID string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error)
	GetSubCPMKByCPMK(cpmkID string) ([]dto.SubCPMKResponse, error)
	UpdateSubCPMK(id string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error)
	DeleteSubCPMK(id string) error

	// Rencana Tugas Methods
	AddRencanaTugas(rpsID string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error)
	GetRencanaTugasByRPS(rpsID string) ([]dto.RPSRencanaTugasResponse, error)
	UpdateRencanaTugas(id string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error)
	DeleteRencanaTugas(id string) error

	// Analisis Ketercapaian CPL Methods
	AddAnalisisKetercapaianCPL(rpsID string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error)
	GetAnalisisKetercapaianCPLByRPS(rpsID string) ([]dto.RPSAnalisisKetercapaianCPLResponse, error)
	UpdateAnalisisKetercapaianCPL(id string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error)
	DeleteAnalisisKetercapaianCPL(id string) error

	// Skala Penilaian Methods
	AddSkalaPenilaian(rpsID string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error)
	GetSkalaPenilaianByRPS(rpsID string) ([]dto.RPSSkalaPenilaianResponse, error)
	UpdateSkalaPenilaian(id string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error)
	DeleteSkalaPenilaian(id string) error
	SetDefaultSkalaPenilaian(rpsID string, req dto.RPSBatchSkalaPenilaianRequest) ([]dto.RPSSkalaPenilaianResponse, error)
}

type rpsExtendedService struct {
	subCpmkService    SubCPMKService
	rencanaTugasSvc   RPSRencanaTugasService
	analisisCPLSvc    RPSAnalisisKetercapaianCPLService
	skalaPenilaianSvc RPSSkalaPenilaianService
}

func NewRPSExtendedService(
	subCpmkService SubCPMKService,
	rencanaTugasSvc RPSRencanaTugasService,
	analisisCPLSvc RPSAnalisisKetercapaianCPLService,
	skalaPenilaianSvc RPSSkalaPenilaianService,
) RPSExtendedService {
	return &rpsExtendedService{
		subCpmkService:    subCpmkService,
		rencanaTugasSvc:   rencanaTugasSvc,
		analisisCPLSvc:    analisisCPLSvc,
		skalaPenilaianSvc: skalaPenilaianSvc,
	}
}

// Sub-CPMK implementations
func (s *rpsExtendedService) AddSubCPMK(cpmkID string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error) {
	return s.subCpmkService.Create(cpmkID, req)
}

func (s *rpsExtendedService) GetSubCPMKByCPMK(cpmkID string) ([]dto.SubCPMKResponse, error) {
	return s.subCpmkService.GetByCPMKID(cpmkID)
}

func (s *rpsExtendedService) UpdateSubCPMK(id string, req dto.SubCPMKRequest) (*dto.SubCPMKResponse, error) {
	return s.subCpmkService.Update(id, req)
}

func (s *rpsExtendedService) DeleteSubCPMK(id string) error {
	return s.subCpmkService.Delete(id)
}

// Rencana Tugas implementations
func (s *rpsExtendedService) AddRencanaTugas(rpsID string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error) {
	return s.rencanaTugasSvc.Create(rpsID, req)
}

func (s *rpsExtendedService) GetRencanaTugasByRPS(rpsID string) ([]dto.RPSRencanaTugasResponse, error) {
	return s.rencanaTugasSvc.GetByRPSID(rpsID)
}

func (s *rpsExtendedService) UpdateRencanaTugas(id string, req dto.RPSRencanaTugasRequest) (*dto.RPSRencanaTugasResponse, error) {
	return s.rencanaTugasSvc.Update(id, req)
}

func (s *rpsExtendedService) DeleteRencanaTugas(id string) error {
	return s.rencanaTugasSvc.Delete(id)
}

// Analisis Ketercapaian CPL implementations
func (s *rpsExtendedService) AddAnalisisKetercapaianCPL(rpsID string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error) {
	return s.analisisCPLSvc.Create(rpsID, req)
}

func (s *rpsExtendedService) GetAnalisisKetercapaianCPLByRPS(rpsID string) ([]dto.RPSAnalisisKetercapaianCPLResponse, error) {
	return s.analisisCPLSvc.GetByRPSID(rpsID)
}

func (s *rpsExtendedService) UpdateAnalisisKetercapaianCPL(id string, req dto.RPSAnalisisKetercapaianCPLRequest) (*dto.RPSAnalisisKetercapaianCPLResponse, error) {
	return s.analisisCPLSvc.Update(id, req)
}

func (s *rpsExtendedService) DeleteAnalisisKetercapaianCPL(id string) error {
	return s.analisisCPLSvc.Delete(id)
}

// Skala Penilaian implementations
func (s *rpsExtendedService) AddSkalaPenilaian(rpsID string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error) {
	return s.skalaPenilaianSvc.Create(rpsID, req)
}

func (s *rpsExtendedService) GetSkalaPenilaianByRPS(rpsID string) ([]dto.RPSSkalaPenilaianResponse, error) {
	return s.skalaPenilaianSvc.GetByRPSID(rpsID)
}

func (s *rpsExtendedService) UpdateSkalaPenilaian(id string, req dto.RPSSkalaPenilaianRequest) (*dto.RPSSkalaPenilaianResponse, error) {
	return s.skalaPenilaianSvc.Update(id, req)
}

func (s *rpsExtendedService) DeleteSkalaPenilaian(id string) error {
	return s.skalaPenilaianSvc.Delete(id)
}

func (s *rpsExtendedService) SetDefaultSkalaPenilaian(rpsID string, req dto.RPSBatchSkalaPenilaianRequest) ([]dto.RPSSkalaPenilaianResponse, error) {
	return s.skalaPenilaianSvc.CreateBatch(rpsID, req)
}
