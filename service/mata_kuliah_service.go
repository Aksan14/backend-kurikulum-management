package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"encoding/json"
	"errors"
)

type MataKuliahService interface {
	GetAllMataKuliah(req dto.MataKuliahListRequest) (*dto.PaginatedResponse, error)
	GetMataKuliahByID(id string) (*dto.MataKuliahResponse, error)
	GetMyMataKuliah(userID string, page, limit int) (*dto.PaginatedResponse, error)
	GetMataKuliahBySemester(semester int, page, limit int) (*dto.PaginatedResponse, error)
	GetMataKuliahByDosen(dosenID string, page, limit int) (*dto.PaginatedResponse, error)
	CreateMataKuliah(userID string, req dto.CreateMataKuliahRequest) (*dto.MataKuliahResponse, error)
	UpdateMataKuliah(id string, req dto.UpdateMataKuliahRequest) (*dto.MataKuliahResponse, error)
	DeleteMataKuliah(id string) error
	ToggleMataKuliahStatus(id string) (*dto.MataKuliahResponse, error)
	AssignDosen(id string, req dto.AssignDosenRequest) (*dto.MataKuliahResponse, error)
	UnassignDosen(id string, dosenType string) (*dto.MataKuliahResponse, error)
}

type mataKuliahService struct {
	mkRepo repository.MataKuliahRepository
}

func NewMataKuliahService(mkRepo repository.MataKuliahRepository) MataKuliahService {
	return &mataKuliahService{mkRepo: mkRepo}
}

func (s *mataKuliahService) GetAllMataKuliah(req dto.MataKuliahListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	mks, total, err := s.mkRepo.FindAll(
		req.Page, req.Limit, req.Search, req.Semester, req.Jenis, req.Status, req.SortBy, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.MataKuliahResponse
	for _, mk := range mks {
		responses = append(responses, toMataKuliahResponse(&mk))
	}

	totalPages := (total + int64(req.Limit) - 1) / int64(req.Limit)

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

func (s *mataKuliahService) GetMataKuliahByID(id string) (*dto.MataKuliahResponse, error) {
	mk, err := s.mkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("mata kuliah tidak ditemukan")
	}

	resp := toMataKuliahResponse(mk)
	return &resp, nil
}

func (s *mataKuliahService) GetMyMataKuliah(userID string, page, limit int) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	mks, total, err := s.mkRepo.FindByDosenID(userID, page, limit)
	if err != nil {
		return nil, err
	}

	var responses []dto.MataKuliahResponse
	for _, mk := range mks {
		responses = append(responses, toMataKuliahResponse(&mk))
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

func (s *mataKuliahService) GetMataKuliahBySemester(semester int, page, limit int) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	mks, total, err := s.mkRepo.FindAll(page, limit, "", semester, "", "aktif", "kode", "asc")
	if err != nil {
		return nil, err
	}

	var responses []dto.MataKuliahResponse
	for _, mk := range mks {
		responses = append(responses, toMataKuliahResponse(&mk))
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

func (s *mataKuliahService) GetMataKuliahByDosen(dosenID string, page, limit int) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	mks, total, err := s.mkRepo.FindByDosenID(dosenID, page, limit)
	if err != nil {
		return nil, err
	}

	var responses []dto.MataKuliahResponse
	for _, mk := range mks {
		responses = append(responses, toMataKuliahResponse(&mk))
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: totalPages,
	}, nil
}

func (s *mataKuliahService) CreateMataKuliah(userID string, req dto.CreateMataKuliahRequest) (*dto.MataKuliahResponse, error) {
	existing, _ := s.mkRepo.FindByKode(req.Kode)
	if existing != nil {
		return nil, errors.New("kode mata kuliah sudah digunakan")
	}

	prasyaratJSON, _ := json.Marshal(req.Prasyarat)

	jenis := req.Jenis
	if jenis == "" {
		jenis = "wajib"
	}

	status := req.Status
	if status == "" {
		status = "aktif"
	}

	mk := &model.MataKuliah{
		Kode:            req.Kode,
		Nama:            req.Nama,
		SKS:             req.SKS,
		Semester:        req.Semester,
		Jenis:           jenis,
		Deskripsi:       req.Deskripsi,
		Prasyarat:       prasyaratJSON,
		Status:          status,
		DosenPengampuID: req.DosenPengampuID,
		KoordinatorID:   req.KoordinatorID,
		IsActive:        true,
		CreatedBy:       userID,
	}

	if err := s.mkRepo.Create(mk); err != nil {
		return nil, errors.New("gagal membuat mata kuliah")
	}

	mk, _ = s.mkRepo.FindByID(mk.ID)
	resp := toMataKuliahResponse(mk)
	return &resp, nil
}

func (s *mataKuliahService) UpdateMataKuliah(id string, req dto.UpdateMataKuliahRequest) (*dto.MataKuliahResponse, error) {
	mk, err := s.mkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("mata kuliah tidak ditemukan")
	}

	if req.Kode != nil && *req.Kode != mk.Kode {
		existing, _ := s.mkRepo.FindByKode(*req.Kode)
		if existing != nil {
			return nil, errors.New("kode mata kuliah sudah digunakan")
		}
		mk.Kode = *req.Kode
	}

	if req.Nama != nil {
		mk.Nama = *req.Nama
	}
	if req.SKS != nil {
		mk.SKS = *req.SKS
	}
	if req.Semester != nil {
		mk.Semester = *req.Semester
	}
	if req.Jenis != nil {
		mk.Jenis = *req.Jenis
	}
	if req.Deskripsi != nil {
		mk.Deskripsi = req.Deskripsi
	}
	if req.Prasyarat != nil {
		prasyaratJSON, _ := json.Marshal(req.Prasyarat)
		mk.Prasyarat = prasyaratJSON
	}
	if req.Status != nil {
		mk.Status = *req.Status
		mk.IsActive = *req.Status == "aktif"
	}
	if req.DosenPengampuID != nil {
		mk.DosenPengampuID = req.DosenPengampuID
	}
	if req.KoordinatorID != nil {
		mk.KoordinatorID = req.KoordinatorID
	}

	if err := s.mkRepo.Update(mk); err != nil {
		return nil, errors.New("gagal mengupdate mata kuliah")
	}

	mk, _ = s.mkRepo.FindByID(mk.ID)
	resp := toMataKuliahResponse(mk)
	return &resp, nil
}

func (s *mataKuliahService) DeleteMataKuliah(id string) error {
	mk, err := s.mkRepo.FindByID(id)
	if err != nil {
		return errors.New("mata kuliah tidak ditemukan")
	}

	mk.Status = "dihapus"
	mk.IsActive = false

	return s.mkRepo.Update(mk)
}

func (s *mataKuliahService) ToggleMataKuliahStatus(id string) (*dto.MataKuliahResponse, error) {
	mk, err := s.mkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("mata kuliah tidak ditemukan")
	}

	if mk.Status == "aktif" {
		mk.Status = "nonaktif"
		mk.IsActive = false
	} else {
		mk.Status = "aktif"
		mk.IsActive = true
	}

	if err := s.mkRepo.Update(mk); err != nil {
		return nil, errors.New("gagal mengubah status mata kuliah")
	}

	resp := toMataKuliahResponse(mk)
	return &resp, nil
}

// AssignDosen assigns dosen pengampu or koordinator to mata kuliah
func (s *mataKuliahService) AssignDosen(id string, req dto.AssignDosenRequest) (*dto.MataKuliahResponse, error) {
	mk, err := s.mkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("mata kuliah tidak ditemukan")
	}

	// Update dosen pengampu if provided
	if req.DosenPengampuID != nil {
		mk.DosenPengampuID = req.DosenPengampuID
	}

	// Update koordinator if provided
	if req.KoordinatorID != nil {
		mk.KoordinatorID = req.KoordinatorID
	}

	if err := s.mkRepo.Update(mk); err != nil {
		return nil, errors.New("gagal assign dosen ke mata kuliah")
	}

	// Reload to get relations
	mk, _ = s.mkRepo.FindByID(mk.ID)

	resp := toMataKuliahResponse(mk)
	return &resp, nil
}

// UnassignDosen removes dosen pengampu or koordinator from mata kuliah
// dosenType: "pengampu", "koordinator", or "all"
func (s *mataKuliahService) UnassignDosen(id string, dosenType string) (*dto.MataKuliahResponse, error) {
	mk, err := s.mkRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("mata kuliah tidak ditemukan")
	}

	switch dosenType {
	case "pengampu":
		if err := s.mkRepo.ClearDosenPengampu(id); err != nil {
			return nil, errors.New("gagal unassign dosen pengampu dari mata kuliah")
		}
	case "koordinator":
		if err := s.mkRepo.ClearKoordinator(id); err != nil {
			return nil, errors.New("gagal unassign koordinator dari mata kuliah")
		}
	case "all":
		if err := s.mkRepo.ClearAllDosen(id); err != nil {
			return nil, errors.New("gagal unassign semua dosen dari mata kuliah")
		}
	default:
		return nil, errors.New("tipe dosen tidak valid. Gunakan: pengampu, koordinator, atau all")
	}

	// Reload to get relations
	mk, _ = s.mkRepo.FindByID(mk.ID)

	resp := toMataKuliahResponse(mk)
	return &resp, nil
}

func toMataKuliahResponse(mk *model.MataKuliah) dto.MataKuliahResponse {
	var prasyarat []string
	if mk.Prasyarat != nil {
		json.Unmarshal(mk.Prasyarat, &prasyarat)
	}
	if prasyarat == nil {
		prasyarat = []string{}
	}

	resp := dto.MataKuliahResponse{
		ID:              mk.ID,
		Kode:            mk.Kode,
		Nama:            mk.Nama,
		SKS:             mk.SKS,
		Semester:        mk.Semester,
		Jenis:           mk.Jenis,
		Deskripsi:       mk.Deskripsi,
		Prasyarat:       prasyarat,
		Status:          mk.Status,
		DosenPengampuID: mk.DosenPengampuID,
		KoordinatorID:   mk.KoordinatorID,
		CreatedAt:       mk.CreatedAt,
		UpdatedAt:       mk.UpdatedAt,
		CreatedBy:       mk.CreatedBy,
	}

	if mk.DosenPengampu != nil && mk.DosenPengampu.ID != "" {
		dosenResp := &dto.UserResponse{
			ID:    mk.DosenPengampu.ID,
			Nama:  mk.DosenPengampu.Nama,
			Email: mk.DosenPengampu.Email,
		}
		resp.DosenPengampu = dosenResp
	}

	if mk.Koordinator != nil && mk.Koordinator.ID != "" {
		koordinatorResp := &dto.UserResponse{
			ID:    mk.Koordinator.ID,
			Nama:  mk.Koordinator.Nama,
			Email: mk.Koordinator.Email,
		}
		resp.Koordinator = koordinatorResp
	}

	return resp
}
