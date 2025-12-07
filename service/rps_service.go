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
	GetAllCPMK() ([]dto.RPSCPMKResponse, error)
	AddCPMK(rpsID string, req dto.RPSCPMKRequest) (*dto.RPSCPMKResponse, error)
	GetCPMKByRPS(rpsID string) ([]dto.RPSCPMKResponse, error)
	UpdateCPMK(id string, req dto.RPSCPMKRequest) (*dto.RPSCPMKResponse, error)
	DeleteCPMK(id string) error

	// Rencana Pembelajaran
	AddRencanaPembelajaran(rpsID string, req dto.RPSRencanaPembelajaranRequest) (*dto.RPSRencanaPembelajaranResponse, error)
	GetRencanaPembelajaranByRPS(rpsID string) ([]dto.RPSRencanaPembelajaranResponse, error)
	UpdateRencanaPembelajaran(id string, req dto.RPSRencanaPembelajaranRequest) (*dto.RPSRencanaPembelajaranResponse, error)
	DeleteRencanaPembelajaran(id string) error

	// Bahan Bacaan
	AddBahanBacaan(rpsID string, req dto.RPSBahanBacaanRequest) (*dto.RPSBahanBacaanResponse, error)
	GetBahanBacaanByRPS(rpsID string) ([]dto.RPSBahanBacaanResponse, error)
	UpdateBahanBacaan(id string, req dto.RPSBahanBacaanRequest) (*dto.RPSBahanBacaanResponse, error)
	DeleteBahanBacaan(id string) error

	// Evaluasi
	AddEvaluasi(rpsID string, req dto.RPSEvaluasiRequest) (*dto.RPSEvaluasiResponse, error)
	GetEvaluasiByRPS(rpsID string) ([]dto.RPSEvaluasiResponse, error)
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
	_ = mk // mk available for future use if needed

	metodePembelajaranJSON, _ := json.Marshal(req.MetodePembelajaran)
	mediaPembelajaranJSON, _ := json.Marshal(req.MediaPembelajaran)

	// Parse tanggal_penyusunan if provided
	var tanggalPenyusunan *time.Time
	if req.TanggalPenyusunan != nil && *req.TanggalPenyusunan != "" {
		parsed, err := time.Parse("2006-01-02", *req.TanggalPenyusunan)
		if err == nil {
			tanggalPenyusunan = &parsed
		}
	}

	rps := &model.RPS{
		MataKuliahID:        req.MataKuliahID,
		TahunAjaran:         req.TahunAjaran,
		SemesterType:        req.SemesterType,
		TanggalPenyusunan:   tanggalPenyusunan,
		DosenID:             dosenID,
		DosenNama:           dosenNama,
		PenyusunNama:        req.PenyusunNama,
		PenyusunNIDN:        req.PenyusunNIDN,
		KoordinatorRMKNama:  req.KoordinatorRMKNama,
		KoordinatorRMKNIDN:  req.KoordinatorRMKNIDN,
		KaprodiNama:         req.KaprodiNama,
		KaprodiNIDN:         req.KaprodiNIDN,
		Fakultas:            req.Fakultas,
		ProgramStudi:        req.ProgramStudi,
		DeskripsiMK:         req.DeskripsiMK,
		CapaianPembelajaran: req.CapaianPembelajaran,
		MetodePembelajaran:  metodePembelajaranJSON,
		MediaPembelajaran:   mediaPembelajaranJSON,
		Status:              "draft",
		Version:             1,
	}

	if err := s.rpsRepo.Create(rps); err != nil {
		return nil, errors.New("gagal membuat RPS")
	}

	rps, _ = s.rpsRepo.FindMinimalByID(rps.ID)
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
		_, err := s.mkRepo.FindByID(req.MataKuliahID)
		if err != nil {
			return nil, errors.New("mata kuliah tidak ditemukan")
		}
		rps.MataKuliahID = req.MataKuliahID
	}

	metodePembelajaranJSON, _ := json.Marshal(req.MetodePembelajaran)
	mediaPembelajaranJSON, _ := json.Marshal(req.MediaPembelajaran)

	// Parse tanggal_penyusunan if provided
	if req.TanggalPenyusunan != nil && *req.TanggalPenyusunan != "" {
		parsed, err := time.Parse("2006-01-02", *req.TanggalPenyusunan)
		if err == nil {
			rps.TanggalPenyusunan = &parsed
		}
	}

	rps.TahunAjaran = req.TahunAjaran
	rps.SemesterType = req.SemesterType
	rps.PenyusunNama = req.PenyusunNama
	rps.PenyusunNIDN = req.PenyusunNIDN
	rps.KoordinatorRMKNama = req.KoordinatorRMKNama
	rps.KoordinatorRMKNIDN = req.KoordinatorRMKNIDN
	rps.KaprodiNama = req.KaprodiNama
	rps.KaprodiNIDN = req.KaprodiNIDN
	rps.Fakultas = req.Fakultas
	rps.ProgramStudi = req.ProgramStudi
	rps.DeskripsiMK = req.DeskripsiMK
	rps.CapaianPembelajaran = req.CapaianPembelajaran
	rps.MetodePembelajaran = metodePembelajaranJSON
	rps.MediaPembelajaran = mediaPembelajaranJSON

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal mengupdate RPS")
	}

	resp := dto.RPSResponse{
		ID:            rps.ID,
		MataKuliahID:  rps.MataKuliahID,
		DosenID:       rps.DosenID,
		Status:        rps.Status,
		ReviewerID:    rps.ReviewerID,
		ReviewCatatan: rps.ReviewCatatan,
		ReviewedAt:    rps.ReviewedAt,
	}
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
	rps.ReviewCatatan = req.ReviewNotes

	switch req.Status {
	case "submitted":
		// submitted status update
	case "approved", "rejected":
		rps.ReviewedAt = &now
		rps.ReviewerID = &reviewerID
	case "published":
		rps.ApprovedAt = &now
	}

	if err := s.rpsRepo.Update(rps); err != nil {
		return nil, errors.New("gagal mengupdate status RPS")
	}

	// Get mata kuliah name for notification
	mkName := ""
	if rps.MataKuliah.ID != "" {
		mkName = rps.MataKuliah.Nama
	}

	// Send notification to dosen
	if s.notifService != nil {
		switch req.Status {
		case "approved":
			relatedID := rps.ID
			s.notifService.Create(dto.CreateNotificationRequest{
				UserID:      rps.DosenID,
				Title:       "RPS Disetujui",
				Message:     "RPS " + mkName + " telah disetujui",
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
				Message:     "RPS " + mkName + " ditolak. Silakan periksa catatan review.",
				Type:        "rejection",
				RelatedID:   &relatedID,
				RelatedType: ptrString("rps"),
				Priority:    "high",
			})
		}
	}

	resp := dto.RPSResponse{
		ID:            rps.ID,
		MataKuliahID:  rps.MataKuliahID,
		DosenID:       rps.DosenID,
		Status:        rps.Status,
		ReviewerID:    rps.ReviewerID,
		ReviewCatatan: rps.ReviewCatatan,
		ReviewedAt:    rps.ReviewedAt,
	}
	return &resp, nil
}

func (s *rpsService) SubmitRPS(id string) (*dto.RPSResponse, error) {
	rps, err := s.rpsRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	// Use UpdateStatus to only update the status field
	if err := s.rpsRepo.UpdateStatus(id, "submitted", "", nil); err != nil {
		return nil, errors.New("gagal submit RPS")
	}

	resp := dto.RPSResponse{
		ID:           rps.ID,
		MataKuliahID: rps.MataKuliahID,
		DosenID:      rps.DosenID,
		Status:       rps.Status,
	}
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
	rps.ApprovedAt = &now
	rps.ReviewerID = &reviewerID
	rps.ReviewCatatan = catatan

	// Use UpdateStatus to only update status-related fields and avoid saving related associations
	if err := s.rpsRepo.UpdateStatus(id, "approved", reviewerID, catatan); err != nil {
		return nil, errors.New("gagal approve RPS")
	}

	// Get mata kuliah name for notification
	mkName := ""
	if rps.MataKuliah.ID != "" {
		mkName = rps.MataKuliah.Nama
	}

	// Send notification
	if s.notifService != nil {
		relatedID := rps.ID
		s.notifService.Create(dto.CreateNotificationRequest{
			UserID:      rps.DosenID,
			Title:       "RPS Disetujui",
			Message:     "RPS " + mkName + " telah disetujui",
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
	rps.ReviewerID = &reviewerID
	rps.ReviewCatatan = alasan

	// Use UpdateStatus to only update status-related fields and avoid saving related associations
	if err := s.rpsRepo.UpdateStatus(id, "rejected", reviewerID, alasan); err != nil {
		return nil, errors.New("gagal reject RPS")
	}

	// Get mata kuliah name for notification
	mkName := ""
	if rps.MataKuliah.ID != "" {
		mkName = rps.MataKuliah.Nama
	}

	// Send notification
	if s.notifService != nil {
		relatedID := rps.ID
		s.notifService.Create(dto.CreateNotificationRequest{
			UserID:      rps.DosenID,
			Title:       "RPS Ditolak",
			Message:     "RPS " + mkName + " ditolak. Silakan periksa catatan review.",
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
	// Use a minimal find to avoid preloading associations that might cause unintended writes
	rps, err := s.rpsRepo.FindMinimalByID(id)
	if err != nil {
		return nil, errors.New("RPS tidak ditemukan")
	}

	now := time.Now()
	rps.Status = "revision"
	rps.ReviewedAt = &now
	rps.ReviewerID = &reviewerID
	rps.ReviewCatatan = catatan

	// Use UpdateStatus to only update status-related fields and avoid saving related associations
	if err := s.rpsRepo.UpdateStatus(id, "revision", reviewerID, catatan); err != nil {
		return nil, errors.New("gagal request revision")
	}

	// Get mata kuliah name for notification
	mkName := ""
	if rps.MataKuliah.ID != "" {
		mkName = rps.MataKuliah.Nama
	}

	// Send notification
	if s.notifService != nil {
		relatedID := rps.ID
		s.notifService.Create(dto.CreateNotificationRequest{
			UserID:      rps.DosenID,
			Title:       "RPS Perlu Revisi",
			Message:     "RPS " + mkName + " perlu direvisi. Silakan periksa catatan review.",
			Type:        "warning",
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
	cpmk := &model.RPSCPMK{
		RPSID:     rpsID,
		Kode:      req.Kode,
		Deskripsi: req.Deskripsi,
		Bobot:     req.Bobot,
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

	cpmk.Kode = req.Kode
	cpmk.Deskripsi = req.Deskripsi
	cpmk.Bobot = req.Bobot
	cpmk.Urutan = req.Urutan

	if err := s.cpmkRepo.Update(cpmk); err != nil {
		return nil, errors.New("gagal mengupdate CPMK")
	}

	resp := toRPSCPMKResponse(cpmk)
	return &resp, nil
}

func (s *rpsService) GetAllCPMK() ([]dto.RPSCPMKResponse, error) {
	cpmkList, err := s.cpmkRepo.FindAll()
	if err != nil {
		return nil, errors.New("gagal mengambil data CPMK")
	}

	var responses []dto.RPSCPMKResponse
	for _, cpmk := range cpmkList {
		responses = append(responses, toRPSCPMKResponse(&cpmk))
	}

	return responses, nil
}

func (s *rpsService) GetCPMKByRPS(rpsID string) ([]dto.RPSCPMKResponse, error) {
	cpmkList, err := s.cpmkRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, errors.New("gagal mengambil data CPMK")
	}

	var responses []dto.RPSCPMKResponse
	for _, cpmk := range cpmkList {
		responses = append(responses, toRPSCPMKResponse(&cpmk))
	}

	return responses, nil
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

func (s *rpsService) GetRencanaPembelajaranByRPS(rpsID string) ([]dto.RPSRencanaPembelajaranResponse, error) {
	rencanaList, err := s.rencanaRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, errors.New("gagal mengambil data rencana pembelajaran")
	}

	var responses []dto.RPSRencanaPembelajaranResponse
	for _, rencana := range rencanaList {
		responses = append(responses, toRPSRencanaPembelajaranResponse(&rencana))
	}

	return responses, nil
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

func (s *rpsService) GetBahanBacaanByRPS(rpsID string) ([]dto.RPSBahanBacaanResponse, error) {
	bahanList, err := s.bahanRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, errors.New("gagal mengambil data bahan bacaan")
	}

	var responses []dto.RPSBahanBacaanResponse
	for _, bahan := range bahanList {
		responses = append(responses, toRPSBahanBacaanResponse(&bahan))
	}

	return responses, nil
}

// Evaluasi methods
func (s *rpsService) AddEvaluasi(rpsID string, req dto.RPSEvaluasiRequest) (*dto.RPSEvaluasiResponse, error) {
	// Check total bobot
	currentTotal, _ := s.evaluasiRepo.GetTotalBobotByRPSID(rpsID)
	if currentTotal+req.Bobot > 100 {
		return nil, errors.New("total bobot evaluasi tidak boleh lebih dari 100%")
	}

	cpmkIDsJSON, _ := json.Marshal(req.CPMKIDs)
	subCPMKIDsJSON, _ := json.Marshal(req.SubCPMKIDs)

	evaluasi := &model.RPSEvaluasi{
		RPSID:             rpsID,
		Komponen:          req.Komponen,
		TeknikPenilaian:   req.TeknikPenilaian,
		Instrumen:         req.Instrumen,
		Bobot:             req.Bobot,
		MingguMulai:       req.MingguMulai,
		MingguSelesai:     req.MingguSelesai,
		CPLID:             req.CPLID,
		KriteriaPenilaian: req.KriteriaPenilaian,
		Urutan:            req.Urutan,
		CPMKIDs:           cpmkIDsJSON,
		SubCPMKIDs:        subCPMKIDsJSON,
		TopikMateri:       req.TopikMateri,
		JenisAssessment:   req.JenisAssessment,
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

	cpmkIDsJSON, _ := json.Marshal(req.CPMKIDs)
	subCPMKIDsJSON, _ := json.Marshal(req.SubCPMKIDs)

	evaluasi.Komponen = req.Komponen
	evaluasi.TeknikPenilaian = req.TeknikPenilaian
	evaluasi.Instrumen = req.Instrumen
	evaluasi.Bobot = req.Bobot
	evaluasi.MingguMulai = req.MingguMulai
	evaluasi.MingguSelesai = req.MingguSelesai
	evaluasi.CPLID = req.CPLID
	evaluasi.KriteriaPenilaian = req.KriteriaPenilaian
	evaluasi.Urutan = req.Urutan
	evaluasi.CPMKIDs = cpmkIDsJSON
	evaluasi.SubCPMKIDs = subCPMKIDsJSON
	evaluasi.TopikMateri = req.TopikMateri
	evaluasi.JenisAssessment = req.JenisAssessment

	if err := s.evaluasiRepo.Update(evaluasi); err != nil {
		return nil, errors.New("gagal mengupdate evaluasi")
	}

	resp := toRPSEvaluasiResponse(evaluasi)
	return &resp, nil
}

func (s *rpsService) DeleteEvaluasi(id string) error {
	return s.evaluasiRepo.Delete(id)
}

func (s *rpsService) GetEvaluasiByRPS(rpsID string) ([]dto.RPSEvaluasiResponse, error) {
	evaluasiList, err := s.evaluasiRepo.FindByRPSID(rpsID)
	if err != nil {
		return nil, errors.New("gagal mengambil data evaluasi")
	}

	var responses []dto.RPSEvaluasiResponse
	for _, evaluasi := range evaluasiList {
		responses = append(responses, toRPSEvaluasiResponse(&evaluasi))
	}

	return responses, nil
}

// Response converters
func toRPSResponse(rps *model.RPS) dto.RPSResponse {
	var metodePembelajaran []string
	var mediaPembelajaran []string
	json.Unmarshal(rps.MetodePembelajaran, &metodePembelajaran)
	json.Unmarshal(rps.MediaPembelajaran, &mediaPembelajaran)

	resp := dto.RPSResponse{
		ID:                  rps.ID,
		MataKuliahID:        rps.MataKuliahID,
		TahunAjaran:         rps.TahunAjaran,
		SemesterType:        rps.SemesterType,
		TanggalPenyusunan:   rps.TanggalPenyusunan,
		DosenID:             rps.DosenID,
		DosenNama:           rps.DosenNama,
		PenyusunID:          rps.PenyusunID,
		PenyusunNama:        rps.PenyusunNama,
		PenyusunNIDN:        rps.PenyusunNIDN,
		KoordinatorRMKID:    rps.KoordinatorRMKID,
		KoordinatorRMKNama:  rps.KoordinatorRMKNama,
		KoordinatorRMKNIDN:  rps.KoordinatorRMKNIDN,
		KaprodiID:           rps.KaprodiID,
		KaprodiNama:         rps.KaprodiNama,
		KaprodiNIDN:         rps.KaprodiNIDN,
		Fakultas:            rps.Fakultas,
		ProgramStudi:        rps.ProgramStudi,
		DeskripsiMK:         rps.DeskripsiMK,
		CapaianPembelajaran: rps.CapaianPembelajaran,
		MetodePembelajaran:  metodePembelajaran,
		MediaPembelajaran:   mediaPembelajaran,
		Status:              rps.Status,
		Version:             rps.Version,
		ReviewerID:          rps.ReviewerID,
		ReviewCatatan:       rps.ReviewCatatan,
		ReviewedAt:          rps.ReviewedAt,
		ApprovedAt:          rps.ApprovedAt,
		CreatedAt:           rps.CreatedAt,
		UpdatedAt:           rps.UpdatedAt,
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

	// Convert Rencana Tugas
	for _, tugas := range rps.RencanaTugas {
		resp.RencanaTugas = append(resp.RencanaTugas, toRPSRencanaTugasResponseFromModel(&tugas))
	}

	// Convert Analisis Ketercapaian
	for _, analisis := range rps.AnalisisKetercapaian {
		resp.AnalisisKetercapaian = append(resp.AnalisisKetercapaian, toRPSAnalisisKetercapaianCPLResponseFromModel(&analisis))
	}

	// Convert Skala Penilaian
	for _, skala := range rps.SkalaPenilaian {
		resp.SkalaPenilaian = append(resp.SkalaPenilaian, toRPSSkalaPenilaianResponseFromModel(&skala))
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
	return dto.RPSCPMKResponse{
		ID:        cpmk.ID,
		RPSID:     cpmk.RPSID,
		Kode:      cpmk.Kode,
		Deskripsi: cpmk.Deskripsi,
		Bobot:     cpmk.Bobot,
		Urutan:    cpmk.Urutan,
		CreatedAt: cpmk.CreatedAt,
		UpdatedAt: cpmk.UpdatedAt,
	}
}

func toRPSRencanaPembelajaranResponse(rencana *model.RPSRencanaPembelajaran) dto.RPSRencanaPembelajaranResponse {
	var subTopik []string
	var cpmkIDs []string
	var subCpmkIDs []string
	var indikator []string
	json.Unmarshal(rencana.SubTopik, &subTopik)
	json.Unmarshal(rencana.CPMKIDs, &cpmkIDs)
	json.Unmarshal(rencana.SubCPMKIDs, &subCpmkIDs)
	json.Unmarshal(rencana.Indikator, &indikator)

	return dto.RPSRencanaPembelajaranResponse{
		ID:                rencana.ID,
		RPSID:             rencana.RPSID,
		Pertemuan:         rencana.Pertemuan,
		MingguMulai:       rencana.MingguMulai,
		MingguSelesai:     rencana.MingguSelesai,
		Topik:             rencana.Topik,
		SubTopik:          subTopik,
		SubCPMKIDs:        subCpmkIDs,
		CPMKIDs:           cpmkIDs,
		Indikator:         indikator,
		Metode:            rencana.Metode,
		MediaLMS:          rencana.MediaLMS,
		Waktu:             rencana.Waktu,
		WaktuTM:           rencana.WaktuTM,
		WaktuBM:           rencana.WaktuBM,
		WaktuPT:           rencana.WaktuPT,
		Materi:            rencana.Materi,
		TeknikPenilaian:   rencana.TeknikPenilaian,
		KriteriaPenilaian: rencana.KriteriaPenilaian,
		BobotPenilaian:    rencana.BobotPenilaian,
		CreatedAt:         rencana.CreatedAt,
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
	var cpmkIDs []string
	var subCPMKIDs []string
	json.Unmarshal(evaluasi.CPMKIDs, &cpmkIDs)
	json.Unmarshal(evaluasi.SubCPMKIDs, &subCPMKIDs)

	resp := dto.RPSEvaluasiResponse{
		ID:                evaluasi.ID,
		RPSID:             evaluasi.RPSID,
		Komponen:          evaluasi.Komponen,
		TeknikPenilaian:   evaluasi.TeknikPenilaian,
		Instrumen:         evaluasi.Instrumen,
		Bobot:             evaluasi.Bobot,
		MingguMulai:       evaluasi.MingguMulai,
		MingguSelesai:     evaluasi.MingguSelesai,
		CPLID:             evaluasi.CPLID,
		KriteriaPenilaian: evaluasi.KriteriaPenilaian,
		Urutan:            evaluasi.Urutan,
		CPMKIDs:           cpmkIDs,
		SubCPMKIDs:        subCPMKIDs,
		TopikMateri:       evaluasi.TopikMateri,
		JenisAssessment:   evaluasi.JenisAssessment,
		CreatedAt:         evaluasi.CreatedAt,
		UpdatedAt:         evaluasi.UpdatedAt,
	}

	if evaluasi.CPL != nil {
		resp.CPL = &dto.CPLSimpleResponse{
			ID:   evaluasi.CPL.ID,
			Kode: evaluasi.CPL.Kode,
			Nama: evaluasi.CPL.Nama,
		}
	}

	return resp
}

func toRPSRencanaTugasResponseFromModel(tugas *model.RPSRencanaTugas) dto.RPSRencanaTugasResponse {
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

func toRPSAnalisisKetercapaianCPLResponseFromModel(analisis *model.RPSAnalisisKetercapaianCPL) dto.RPSAnalisisKetercapaianCPLResponse {
	var cpmkIDs []string
	var subCPMKIDs []string
	json.Unmarshal(analisis.CPMKIDs, &cpmkIDs)
	json.Unmarshal(analisis.SubCPMKIDs, &subCPMKIDs)

	return dto.RPSAnalisisKetercapaianCPLResponse{
		ID:              analisis.ID,
		RPSID:           analisis.RPSID,
		MingguMulai:     analisis.MingguMulai,
		MingguSelesai:   analisis.MingguSelesai,
		CPLID:           analisis.CPLID,
		CPMKIDs:         cpmkIDs,
		SubCPMKIDs:      subCPMKIDs,
		TopikMateri:     analisis.TopikMateri,
		JenisAssessment: analisis.JenisAssessment,
		BobotKontribusi: analisis.BobotKontribusi,
		CreatedAt:       analisis.CreatedAt,
		UpdatedAt:       analisis.UpdatedAt,
	}
}

func toRPSSkalaPenilaianResponseFromModel(skala *model.RPSSkalaPenilaian) dto.RPSSkalaPenilaianResponse {
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
