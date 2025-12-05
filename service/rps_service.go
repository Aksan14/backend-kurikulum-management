package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"encoding/json"
	"errors"
	"time"
)

type RPSService interface {
	GetAllRPS(req dto.RPSListRequest) (*dto.PaginatedResponse, error)
	GetRPSByID(id string) (*dto.RPSResponse, error)
	GetRPSByDosenID(dosenID string, page, limit int, status string) (*dto.PaginatedResponse, error)
	GetRPSByDosen(dosenID string, status string) ([]dto.RPSResponse, error)
	GetRPSByMataKuliah(mkID string) ([]dto.RPSResponse, error)
	CreateRPS(dosenID, dosenNama string, req dto.RPSRequest) (*dto.RPSResponse, error)
	UpdateRPS(id string, req dto.RPSRequest) (*dto.RPSResponse, error)
	DeleteRPS(id string) error
	UpdateRPSStatus(id, reviewerID string, req dto.RPSStatusUpdateRequest) (*dto.RPSResponse, error)
	SubmitRPS(id string) (*dto.RPSResponse, error)
	ApproveRPS(id, reviewerID string, catatan *string) (*dto.RPSResponse, error)
	RejectRPS(id, reviewerID string, alasan *string) (*dto.RPSResponse, error)
	RequestRevision(id, reviewerID string, catatan *string) (*dto.RPSResponse, error)

	// CPMK
	AddCPMK(rpsID string, req dto.RPSCPMKRequest) (*dto.RPSCPMKResponse, error)
	UpdateCPMK(id string, req dto.RPSCPMKRequest) (*dto.RPSCPMKResponse, error)
	DeleteCPMK(id string) error

	// Rencana Pembelajaran
	AddRencanaPembelajaran(rpsID string, req dto.RPSRencanaPembelajaranRequest) (*dto.RPSRencanaPembelajaranResponse, error)
	UpdateRencanaPembelajaran(id string, req dto.RPSRencanaPembelajaranRequest) (*dto.RPSRencanaPembelajaranResponse, error)
	DeleteRencanaPembelajaran(id string) error

	// Bahan Bacaan
	AddBahanBacaan(rpsID string, req dto.RPSBahanBacaanRequest) (*dto.RPSBahanBacaanResponse, error)
	UpdateBahanBacaan(id string, req dto.RPSBahanBacaanRequest) (*dto.RPSBahanBacaanResponse, error)
	DeleteBahanBacaan(id string) error

	// Evaluasi
	AddEvaluasi(rpsID string, req dto.RPSEvaluasiRequest) (*dto.RPSEvaluasiResponse, error)
	UpdateEvaluasi(id string, req dto.RPSEvaluasiRequest) (*dto.RPSEvaluasiResponse, error)
	DeleteEvaluasi(id string) error
}

type rpsService struct {
	rpsRepo      repository.RPSRepository
	mkRepo       repository.MataKuliahRepository
	cpmkRepo     repository.RPSCPMKRepository
	rencanaRepo  repository.RPSRencanaPembelajaranRepository
	bahanRepo    repository.RPSBahanBacaanRepository
	evaluasiRepo repository.RPSEvaluasiRepository
	notifService NotificationService
}

func NewRPSService(
	rpsRepo repository.RPSRepository,
	mkRepo repository.MataKuliahRepository,
	cpmkRepo repository.RPSCPMKRepository,
	rencanaRepo repository.RPSRencanaPembelajaranRepository,
	bahanRepo repository.RPSBahanBacaanRepository,
	evaluasiRepo repository.RPSEvaluasiRepository,
	notifService NotificationService,
) RPSService {
	return &rpsService{
		rpsRepo:      rpsRepo,
		mkRepo:       mkRepo,
		cpmkRepo:     cpmkRepo,
		rencanaRepo:  rencanaRepo,
		bahanRepo:    bahanRepo,
		evaluasiRepo: evaluasiRepo,
		notifService: notifService,
	}
}

func (s *rpsService) GetAllRPS(req dto.RPSListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	rpsList, total, err := s.rpsRepo.FindAll(
		req.Page, req.Limit, req.Search, req.MataKuliahID, req.DosenID,
		req.Status, req.TahunAkademik, req.Semester, req.SortBy, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSResponse
	for _, rps := range rpsList {
		responses = append(responses, toRPSResponse(&rps))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *rpsService) GetRPSByID(id string) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) GetRPSByDosenID(dosenID string, page, limit int, status string) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	rpsList, total, err := s.rpsRepo.FindByDosenID(dosenID, page, limit, status)
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSResponse
	for _, rps := range rpsList {
		responses = append(responses, toRPSResponse(&rps))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: (total + int64(limit) - 1) / int64(limit),
	}, nil
}

func (s *rpsService) CreateRPS(dosenID, dosenNama string, req dto.RPSRequest) (*dto.RPSResponse, error) {
	// Get mata kuliah info
	mk, err := s.mkRepo.FindByID(req.MataKuliahID)
	if err != nil {
		return nil, errors.New("mata kuliah tidak ditemukan")
	}

	metodeJSON, _ := json.Marshal(req.Metode)
	bobotJSON, _ := json.Marshal(req.BobotNilai)

	rps := &model.RPS{
		MataKuliahID:   req.MataKuliahID,
		MataKuliahNama: mk.Nama,
		KodeMK:         mk.Kode,
		SKS:            mk.SKS,
		Semester:       mk.Semester,
		TahunAkademik:  req.TahunAkademik,
		DosenID:        dosenID,
		DosenNama:      dosenNama,
		Deskripsi:      req.Deskripsi,
		Tujuan:         req.Tujuan,
		Metode:         metodeJSON,
		BobotNilai:     bobotJSON,
		Status:         "draft",
	}

	if err := s.rpsRepo.Create(rps); err != nil {
		return nil, errors.New("gagal membuat RPS")
	}

	rps, _ = s.rpsRepo.FindByID(rps.ID)
	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) UpdateRPS(id string, req dto.RPSRequest) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	if rps.Status == "approved" || rps.Status == "published" {
		return nil, errors.New("RPS yang sudah disetujui tidak dapat diubah")
	}

	// Get mata kuliah info if changed
	if req.MataKuliahID != rps.MataKuliahID {
		mk, err := s.mkRepo.FindByID(req.MataKuliahID)
		if err != nil {
			return nil, errors.New("mata kuliah tidak ditemukan")
		}
		rps.MataKuliahID = req.MataKuliahID
		rps.MataKuliahNama = mk.Nama
		rps.KodeMK = mk.Kode
		rps.SKS = mk.SKS
		rps.Semester = mk.Semester
	}

	metodeJSON, _ := json.Marshal(req.Metode)
	bobotJSON, _ := json.Marshal(req.BobotNilai)

	rps.TahunAkademik = req.TahunAkademik
	rps.Deskripsi = req.Deskripsi
	rps.Tujuan = req.Tujuan
	rps.Metode = metodeJSON
	rps.BobotNilai = bobotJSON

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal mengupdate RPS")
	}

	rps, _ = s.rpsRepo.FindByID(rps.ID)
	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) DeleteRPS(id string) error {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return errors.New("RPS tidak ditemukan")
	}

	if rps.Status == "approved" || rps.Status == "published" {
		return errors.New("RPS yang sudah disetujui tidak dapat dihapus")
	}

	return s.rpsRepo.Delete(id)
}

func (s *rpsService) UpdateRPSStatus(id, reviewerID string, req dto.RPSStatusUpdateRequest) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	now := time.Now()
	rps.Status = req.Status
	rps.ReviewNotes = req.ReviewNotes

	switch req.Status {
	case "submitted":
		rps.SubmittedAt = &now
	case "approved", "rejected":
		rps.ReviewedAt = &now
		rps.ReviewedBy = &reviewerID
	case "published":
		rps.PublishedAt = &now
	}

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal mengupdate status RPS")
	}

	// Send notification to dosen
	if s.notifService != nil {
		switch req.Status {
		case "approved":
			relatedID := rps.ID
			s.notifService.Create(dto.CreateNotificationRequest{
				UserID:      rps.DosenID,
				Title:       "RPS Disetujui",
				Message:     "RPS " + rps.MataKuliahNama + " telah disetujui",
				Type:        "approval",
				RelatedID:   &relatedID,
				RelatedType: ptrString("rps"),
				Priority:    "normal",
			})
		case "rejected":
			relatedID := rps.ID
			s.notifService.Create(dto.CreateNotificationRequest{
				UserID:      rps.DosenID,
				Title:       "RPS Ditolak",
				Message:     "RPS " + rps.MataKuliahNama + " ditolak. Silakan periksa catatan review.",
				Type:        "rejection",
				RelatedID:   &relatedID,
				RelatedType: ptrString("rps"),
				Priority:    "high",
			})
		}
	}

	rps, _ = s.rpsRepo.FindByID(rps.ID)
	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) SubmitRPS(id string) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	now := time.Now()
	rps.Status = "submitted"
	rps.SubmittedAt = &now

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal submit RPS")
	}

	rps, _ = s.rpsRepo.FindByID(rps.ID)
	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) ApproveRPS(id, reviewerID string, catatan *string) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	now := time.Now()
	rps.Status = "approved"
	rps.ReviewedAt = &now
	rps.ReviewedBy = &reviewerID
	rps.ReviewNotes = catatan

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal approve RPS")
	}

	// Send notification
	if s.notifService != nil {
		relatedID := rps.ID
		s.notifService.Create(dto.CreateNotificationRequest{
			UserID:      rps.DosenID,
			Title:       "RPS Disetujui",
			Message:     "RPS " + rps.MataKuliahNama + " telah disetujui",
			Type:        "approval",
			RelatedID:   &relatedID,
			RelatedType: ptrString("rps"),
			Priority:    "normal",
		})
	}

	rps, _ = s.rpsRepo.FindByID(rps.ID)
	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) RejectRPS(id, reviewerID string, alasan *string) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	now := time.Now()
	rps.Status = "rejected"
	rps.ReviewedAt = &now
	rps.ReviewedBy = &reviewerID
	rps.ReviewNotes = alasan

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal reject RPS")
	}

	// Send notification
	if s.notifService != nil {
		relatedID := rps.ID
		s.notifService.Create(dto.CreateNotificationRequest{
			UserID:      rps.DosenID,
			Title:       "RPS Ditolak",
			Message:     "RPS " + rps.MataKuliahNama + " ditolak. Silakan periksa catatan review.",
			Type:        "rejection",
			RelatedID:   &relatedID,
			RelatedType: ptrString("rps"),
			Priority:    "high",
		})
	}

	rps, _ = s.rpsRepo.FindByID(rps.ID)
	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) RequestRevision(id, reviewerID string, catatan *string) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	now := time.Now()
	rps.Status = "revision_requested"
	rps.ReviewedAt = &now
	rps.ReviewedBy = &reviewerID
	rps.ReviewNotes = catatan

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal request revision")
	}

	// Send notification
	if s.notifService != nil {
		relatedID := rps.ID
		s.notifService.Create(dto.CreateNotificationRequest{
			UserID:      rps.DosenID,
			Title:       "RPS Perlu Revisi",
			Message:     "RPS " + rps.MataKuliahNama + " perlu direvisi. Silakan periksa catatan review.",
			Type:        "revision",
			RelatedID:   &relatedID,
			RelatedType: ptrString("rps"),
			Priority:    "high",
		})
	}

	rps, _ = s.rpsRepo.FindByID(rps.ID)
	resp := toRPSResponse(rps)
	return &resp, nil
}

func (s *rpsService) GetRPSByDosen(dosenID string, status string) ([]dto.RPSResponse, error) {
	rpsList, _, err := s.rpsRepo.FindByDosenID(dosenID, 1, 1000, status)
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSResponse
	for _, rps := range rpsList {
		responses = append(responses, toRPSResponse(&rps))
	}

	return responses, nil
}

func (s *rpsService) GetRPSByMataKuliah(mkID string) ([]dto.RPSResponse, error) {
	rpsList, _, err := s.rpsRepo.FindAll(1, 1000, "", mkID, "", "", "", 0, "", "")
	if err != nil {
		return nil, err
	}

	var responses []dto.RPSResponse
	for _, rps := range rpsList {
		responses = append(responses, toRPSResponse(&rps))
	}

	return responses, nil
}

// CPMK methods
func (s *rpsService) AddCPMK(rpsID string, req dto.RPSCPMKRequest) (*dto.RPSCPMKResponse, error) {
	cplIDsJSON, _ := json.Marshal(req.CPLIDs)

	cpmk := &model.RPSCPMK{
		RPSID:     rpsID,
		Kode:      req.Kode,
		Deskripsi: req.Deskripsi,
		CPLIDs:    cplIDsJSON,
		Urutan:    req.Urutan,
	}

	if err := s.cpmkRepo.Create(cpmk); err != nil {
		return nil, errors.New("gagal menambah CPMK")
	}

	resp := toRPSCPMKResponse(cpmk)
	return &resp, nil
}

func (s *rpsService) UpdateCPMK(id string, req dto.RPSCPMKRequest) (*dto.RPSCPMKResponse, error) {
	cpmk, err := s.cpmkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("CPMK tidak ditemukan")
	}

	cplIDsJSON, _ := json.Marshal(req.CPLIDs)

	cpmk.Kode = req.Kode
	cpmk.Deskripsi = req.Deskripsi
	cpmk.CPLIDs = cplIDsJSON
	cpmk.Urutan = req.Urutan

	if err := s.cpmkRepo.Update(cpmk); err != nil {
		return nil, errors.New("gagal mengupdate CPMK")
	}

	resp := toRPSCPMKResponse(cpmk)
	return &resp, nil
}

func (s *rpsService) DeleteCPMK(id string) error {
	return s.cpmkRepo.Delete(id)
}

// Rencana Pembelajaran methods
func (s *rpsService) AddRencanaPembelajaran(rpsID string, req dto.RPSRencanaPembelajaranRequest) (*dto.RPSRencanaPembelajaranResponse, error) {
	subTopikJSON, _ := json.Marshal(req.SubTopik)
	cpmkIDsJSON, _ := json.Marshal(req.CPMKIDs)

	rencana := &model.RPSRencanaPembelajaran{
		RPSID:     rpsID,
		Pertemuan: req.Pertemuan,
		Topik:     req.Topik,
		SubTopik:  subTopikJSON,
		Metode:    req.Metode,
		Waktu:     req.Waktu,
		CPMKIDs:   cpmkIDsJSON,
		Materi:    req.Materi,
	}

	if err := s.rencanaRepo.Create(rencana); err != nil {
		return nil, errors.New("gagal menambah rencana pembelajaran")
	}

	resp := toRPSRencanaPembelajaranResponse(rencana)
	return &resp, nil
}

func (s *rpsService) UpdateRencanaPembelajaran(id string, req dto.RPSRencanaPembelajaranRequest) (*dto.RPSRencanaPembelajaranResponse, error) {
	rencana, err := s.rencanaRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("rencana pembelajaran tidak ditemukan")
	}

	subTopikJSON, _ := json.Marshal(req.SubTopik)
	cpmkIDsJSON, _ := json.Marshal(req.CPMKIDs)

	rencana.Pertemuan = req.Pertemuan
	rencana.Topik = req.Topik
	rencana.SubTopik = subTopikJSON
	rencana.Metode = req.Metode
	rencana.Waktu = req.Waktu
	rencana.CPMKIDs = cpmkIDsJSON
	rencana.Materi = req.Materi

	if err := s.rencanaRepo.Update(rencana); err != nil {
		return nil, errors.New("gagal mengupdate rencana pembelajaran")
	}

	resp := toRPSRencanaPembelajaranResponse(rencana)
	return &resp, nil
}

func (s *rpsService) DeleteRencanaPembelajaran(id string) error {
	return s.rencanaRepo.Delete(id)
}

// Bahan Bacaan methods
func (s *rpsService) AddBahanBacaan(rpsID string, req dto.RPSBahanBacaanRequest) (*dto.RPSBahanBacaanResponse, error) {
	bahan := &model.RPSBahanBacaan{
		RPSID:   rpsID,
		Judul:   req.Judul,
		Penulis: req.Penulis,
		Tahun:   req.Tahun,
		Jenis:   req.Jenis,
		URL:     req.URL,
		ISBN:    req.ISBN,
		Halaman: req.Halaman,
		Urutan:  req.Urutan,
	}

	if err := s.bahanRepo.Create(bahan); err != nil {
		return nil, errors.New("gagal menambah bahan bacaan")
	}

	resp := toRPSBahanBacaanResponse(bahan)
	return &resp, nil
}

func (s *rpsService) UpdateBahanBacaan(id string, req dto.RPSBahanBacaanRequest) (*dto.RPSBahanBacaanResponse, error) {
	bahan, err := s.bahanRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("bahan bacaan tidak ditemukan")
	}

	bahan.Judul = req.Judul
	bahan.Penulis = req.Penulis
	bahan.Tahun = req.Tahun
	bahan.Jenis = req.Jenis
	bahan.URL = req.URL
	bahan.ISBN = req.ISBN
	bahan.Halaman = req.Halaman
	bahan.Urutan = req.Urutan

	if err := s.bahanRepo.Update(bahan); err != nil {
		return nil, errors.New("gagal mengupdate bahan bacaan")
	}

	resp := toRPSBahanBacaanResponse(bahan)
	return &resp, nil
}

func (s *rpsService) DeleteBahanBacaan(id string) error {
	return s.bahanRepo.Delete(id)
}

// Evaluasi methods
func (s *rpsService) AddEvaluasi(rpsID string, req dto.RPSEvaluasiRequest) (*dto.RPSEvaluasiResponse, error) {
	// Check total bobot
	currentTotal, _ := s.evaluasiRepo.GetTotalBobotByRPSID(rpsID)
	if currentTotal+req.Bobot > 100 {
		return nil, errors.New("total bobot evaluasi tidak boleh lebih dari 100%")
	}

	mingguJSON, _ := json.Marshal(req.MingguPelaksanaan)

	evaluasi := &model.RPSEvaluasi{
		RPSID:             rpsID,
		Jenis:             req.Jenis,
		Bobot:             req.Bobot,
		Deskripsi:         req.Deskripsi,
		MingguPelaksanaan: mingguJSON,
		KriteriaPenilaian: req.KriteriaPenilaian,
		RubrikPenilaian:   req.RubrikPenilaian,
	}

	if err := s.evaluasiRepo.Create(evaluasi); err != nil {
		return nil, errors.New("gagal menambah evaluasi")
	}

	resp := toRPSEvaluasiResponse(evaluasi)
	return &resp, nil
}

func (s *rpsService) UpdateEvaluasi(id string, req dto.RPSEvaluasiRequest) (*dto.RPSEvaluasiResponse, error) {
	evaluasi, err := s.evaluasiRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("evaluasi tidak ditemukan")
	}

	// Check total bobot (excluding current)
	currentTotal, _ := s.evaluasiRepo.GetTotalBobotByRPSID(evaluasi.RPSID)
	if (currentTotal-evaluasi.Bobot)+req.Bobot > 100 {
		return nil, errors.New("total bobot evaluasi tidak boleh lebih dari 100%")
	}

	mingguJSON, _ := json.Marshal(req.MingguPelaksanaan)

	evaluasi.Jenis = req.Jenis
	evaluasi.Bobot = req.Bobot
	evaluasi.Deskripsi = req.Deskripsi
	evaluasi.MingguPelaksanaan = mingguJSON
	evaluasi.KriteriaPenilaian = req.KriteriaPenilaian
	evaluasi.RubrikPenilaian = req.RubrikPenilaian

	if err := s.evaluasiRepo.Update(evaluasi); err != nil {
		return nil, errors.New("gagal mengupdate evaluasi")
	}

	resp := toRPSEvaluasiResponse(evaluasi)
	return &resp, nil
}

func (s *rpsService) DeleteEvaluasi(id string) error {
	return s.evaluasiRepo.Delete(id)
}

// Response converters
func toRPSResponse(rps *model.RPS) dto.RPSResponse {
	var metode []string
	var bobotNilai dto.BobotNilaiRequest
	json.Unmarshal(rps.Metode, &metode)
	json.Unmarshal(rps.BobotNilai, &bobotNilai)

	resp := dto.RPSResponse{
		ID:             rps.ID,
		MataKuliahID:   rps.MataKuliahID,
		MataKuliahNama: rps.MataKuliahNama,
		KodeMK:         rps.KodeMK,
		SKS:            rps.SKS,
		Semester:       rps.Semester,
		TahunAkademik:  rps.TahunAkademik,
		DosenID:        rps.DosenID,
		DosenNama:      rps.DosenNama,
		Deskripsi:      rps.Deskripsi,
		Tujuan:         rps.Tujuan,
		Metode:         metode,
		BobotNilai:     bobotNilai,
		Status:         rps.Status,
		CreatedAt:      rps.CreatedAt,
		UpdatedAt:      rps.UpdatedAt,
		SubmittedAt:    rps.SubmittedAt,
		ReviewedAt:     rps.ReviewedAt,
		PublishedAt:    rps.PublishedAt,
		ReviewedBy:     rps.ReviewedBy,
		ReviewNotes:    rps.ReviewNotes,
	}

	// Convert CPMK
	for _, cpmk := range rps.CPMK {
		resp.CPMK = append(resp.CPMK, toRPSCPMKResponse(&cpmk))
	}

	// Convert Rencana Pembelajaran
	for _, rencana := range rps.RencanaPembelajaran {
		resp.RencanaPembelajaran = append(resp.RencanaPembelajaran, toRPSRencanaPembelajaranResponse(&rencana))
	}

	// Convert Bahan Bacaan
	for _, bahan := range rps.BahanBacaan {
		resp.BahanBacaan = append(resp.BahanBacaan, toRPSBahanBacaanResponse(&bahan))
	}

	// Convert Evaluasi
	for _, eval := range rps.Evaluasi {
		resp.Evaluasi = append(resp.Evaluasi, toRPSEvaluasiResponse(&eval))
	}

	if rps.Dosen.ID != "" {
		dosen := toUserResponse(&rps.Dosen)
		resp.Dosen = &dosen
	}

	if rps.MataKuliah.ID != "" {
		mk := toMataKuliahResponse(&rps.MataKuliah)
		resp.MataKuliah = &mk
	}

	return resp
}

func toRPSCPMKResponse(cpmk *model.RPSCPMK) dto.RPSCPMKResponse {
	var cplIDs []string
	json.Unmarshal(cpmk.CPLIDs, &cplIDs)

	return dto.RPSCPMKResponse{
		ID:        cpmk.ID,
		RPSID:     cpmk.RPSID,
		Kode:      cpmk.Kode,
		Deskripsi: cpmk.Deskripsi,
		CPLIDs:    cplIDs,
		Urutan:    cpmk.Urutan,
		CreatedAt: cpmk.CreatedAt,
	}
}

func toRPSRencanaPembelajaranResponse(rencana *model.RPSRencanaPembelajaran) dto.RPSRencanaPembelajaranResponse {
	var subTopik []string
	var cpmkIDs []string
	json.Unmarshal(rencana.SubTopik, &subTopik)
	json.Unmarshal(rencana.CPMKIDs, &cpmkIDs)

	return dto.RPSRencanaPembelajaranResponse{
		ID:        rencana.ID,
		RPSID:     rencana.RPSID,
		Pertemuan: rencana.Pertemuan,
		Topik:     rencana.Topik,
		SubTopik:  subTopik,
		Metode:    rencana.Metode,
		Waktu:     rencana.Waktu,
		CPMKIDs:   cpmkIDs,
		Materi:    rencana.Materi,
		CreatedAt: rencana.CreatedAt,
	}
}

func toRPSBahanBacaanResponse(bahan *model.RPSBahanBacaan) dto.RPSBahanBacaanResponse {
	return dto.RPSBahanBacaanResponse{
		ID:        bahan.ID,
		RPSID:     bahan.RPSID,
		Judul:     bahan.Judul,
		Penulis:   bahan.Penulis,
		Tahun:     bahan.Tahun,
		Jenis:     bahan.Jenis,
		URL:       bahan.URL,
		ISBN:      bahan.ISBN,
		Halaman:   bahan.Halaman,
		Urutan:    bahan.Urutan,
		CreatedAt: bahan.CreatedAt,
	}
}

func toRPSEvaluasiResponse(evaluasi *model.RPSEvaluasi) dto.RPSEvaluasiResponse {
	var minggu []int
	json.Unmarshal(evaluasi.MingguPelaksanaan, &minggu)

	return dto.RPSEvaluasiResponse{
		ID:                evaluasi.ID,
		RPSID:             evaluasi.RPSID,
		Jenis:             evaluasi.Jenis,
		Bobot:             evaluasi.Bobot,
		Deskripsi:         evaluasi.Deskripsi,
		MingguPelaksanaan: minggu,
		KriteriaPenilaian: evaluasi.KriteriaPenilaian,
		RubrikPenilaian:   evaluasi.RubrikPenilaian,
		CreatedAt:         evaluasi.CreatedAt,
	}
}
