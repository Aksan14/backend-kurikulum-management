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
	tugasRepo   repository.RPSRencanaTugasRepository
	rpsRepo     repository.RPSRepository
	subCpmkRepo repository.SubCPMKRepository
}

func NewRPSRencanaTugasService(
	tugasRepo repository.RPSRencanaTugasRepository,
	rpsRepo repository.RPSRepository,
	subCpmkRepo repository.SubCPMKRepository,
) RPSRencanaTugasService {
	return &rpsRencanaTugasService{
		tugasRepo:   tugasRepo,
		rpsRepo:     rpsRepo,
		subCpmkRepo: subCpmkRepo,
	}
}

func (s *rpsRencanaTugasService) GetByRPSID(rpsID string) ([]dto.RPSRencanaTugasResponse, error) {
	tugasList, err := s.tugasRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSRencanaTugasResponse
	for _, tugas := range tugasList {
		responses = append(responses, s.toRPSRencanaTugasResponseWithSubCPMK(&tugas))
	}
	return responses, nil
}

func (s *rpsRencanaTugasService) GetByID(id string) (*dto.RPSRencanaTugasResponse, error) {
	tugas, err := s.tugasRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("rencana tugas tidak ditemukan")
	}
	resp := s.toRPSRencanaTugasResponseWithSubCPMK(tugas)
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
		SubCPMKID:             req.SubCPMKID,
		DaftarRujukan:         req.DaftarRujukan,
	}

	if err := s.tugasRepo.Create(tugas); err != nil {
		return nil, errors.New("gagal membuat rencana tugas")
	}

	resp := s.toRPSRencanaTugasResponseWithSubCPMK(tugas)
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
	tugas.SubCPMKID = req.SubCPMKID
	tugas.DaftarRujukan = req.DaftarRujukan

	if err := s.tugasRepo.Update(tugas); err != nil {
		return nil, errors.New("gagal mengupdate rencana tugas")
	}

	resp := s.toRPSRencanaTugasResponseWithSubCPMK(tugas)
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
		SubCPMKID:             tugas.SubCPMKID,
		DaftarRujukan:         tugas.DaftarRujukan,
		CreatedAt:             tugas.CreatedAt,
		UpdatedAt:             tugas.UpdatedAt,
	}
}

// toRPSRencanaTugasResponseWithSubCPMK - mengambil data SubCPMK berdasarkan SubCPMKID (Indikator = deskripsi sub_cpmk)
func (s *rpsRencanaTugasService) toRPSRencanaTugasResponseWithSubCPMK(tugas *model.RPSRencanaTugas) dto.RPSRencanaTugasResponse {
	resp := toRPSRencanaTugasResponse(tugas)

	// Populate SubCPMK dari database berdasarkan SubCPMKID
	if resp.SubCPMKID != nil && *resp.SubCPMKID != "" {
		subCpmk, err := s.subCpmkRepo.FindByID(*resp.SubCPMKID)
		if err == nil && subCpmk != nil {
			resp.SubCPMK = &dto.SubCPMKResponse{
				ID:        subCpmk.ID,
				CPMKID:    subCpmk.CPMKID,
				Kode:      subCpmk.Kode,
				Deskripsi: subCpmk.Deskripsi, // Ini adalah INDIKATOR
				Urutan:    subCpmk.Urutan,
				CreatedAt: subCpmk.CreatedAt,
				UpdatedAt: subCpmk.UpdatedAt,
			}
		}
	}

	return resp
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
}

type rpsExtendedService struct {
	subCpmkService  SubCPMKService
	rencanaTugasSvc RPSRencanaTugasService
	analisisCPLSvc  RPSAnalisisKetercapaianCPLService
}

func NewRPSExtendedService(
	subCpmkService SubCPMKService,
	rencanaTugasSvc RPSRencanaTugasService,
	analisisCPLSvc RPSAnalisisKetercapaianCPLService,
) RPSExtendedService {
	return &rpsExtendedService{
		subCpmkService:  subCpmkService,
		rencanaTugasSvc: rencanaTugasSvc,
		analisisCPLSvc:  analisisCPLSvc,
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
