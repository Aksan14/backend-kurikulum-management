package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"errors"
)

type CPLService interface {
	GetAllCPL(req dto.CPLListRequest) (*dto.PaginatedResponse, error)
	GetCPLByID(id string) (*dto.CPLResponse, error)
	CreateCPL(userID string, req dto.CreateCPLRequest) (*dto.CPLResponse, error)
	UpdateCPL(id string, req dto.UpdateCPLRequest) (*dto.CPLResponse, error)
	DeleteCPL(id string) error
	UpdateStatus(id string, req dto.CPLStatusUpdateRequest) (*dto.CPLResponse, error)
	UpdateCPLStatus(id string, status string) (*dto.CPLResponse, error)
	GetCPLStatistics() (*dto.CPLStatisticsResponse, error)
	GetActiveCPL() ([]dto.CPLResponse, error)
}

type cplService struct {
	cplRepo repository.CPLRepository
}

func NewCPLService(cplRepo repository.CPLRepository) CPLService {
	return &cplService{cplRepo: cplRepo}
}

func (s *cplService) GetAllCPL(req dto.CPLListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	cpls, total, err := s.cplRepo.FindAll(
		req.Page, req.Limit, req.Search, req.Status, "", "", req.SortBy, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.CPLResponse
	for _, cpl := range cpls {
		responses = append(responses, toCPLResponse(&cpl))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *cplService) GetCPLByID(id string) (*dto.CPLResponse, error) {
	cpl, err := s.cplRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("CPL tidak ditemukan")
	}

	resp := toCPLResponse(cpl)
	return &resp, nil
}

func (s *cplService) CreateCPL(userID string, req dto.CreateCPLRequest) (*dto.CPLResponse, error) {
	// Check duplicate kode
	existing, _ := s.cplRepo.FindByKode(req.Kode)
	if existing != nil {
		return nil, errors.New("kode CPL sudah digunakan")
	}

	cpl := &model.CPL{
		Kode:      req.Kode,
		Nama:      req.Nama,
		Deskripsi: req.Deskripsi,
		Status:    "draft",
		Version:   1,
		CreatedBy: userID,
	}

	if err := s.cplRepo.Create(cpl); err != nil {
		return nil, errors.New("gagal membuat CPL")
	}

	resp := toCPLResponse(cpl)
	return &resp, nil
}

func (s *cplService) UpdateCPL(id string, req dto.UpdateCPLRequest) (*dto.CPLResponse, error) {
	cpl, err := s.cplRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("CPL tidak ditemukan")
	}

	// Check duplicate kode if changed
	if req.Kode != "" && req.Kode != cpl.Kode {
		existing, _ := s.cplRepo.FindByKode(req.Kode)
		if existing != nil {
			return nil, errors.New("kode CPL sudah digunakan")
		}
		cpl.Kode = req.Kode
	}

	if req.Nama != "" {
		cpl.Nama = req.Nama
	}
	if req.Deskripsi != nil {
		cpl.Deskripsi = req.Deskripsi
	}
	if req.Status != "" {
		cpl.Status = req.Status
	}
	cpl.Version++

	if err := s.cplRepo.Update(cpl); err != nil {
		return nil, errors.New("gagal mengupdate CPL")
	}

	resp := toCPLResponse(cpl)
	return &resp, nil
}

func (s *cplService) DeleteCPL(id string) error {
	_, err := s.cplRepo.FindByID(id)
	if err != nil {
		return errors.New("CPL tidak ditemukan")
	}

	return s.cplRepo.Delete(id)
}

func (s *cplService) UpdateStatus(id string, req dto.CPLStatusUpdateRequest) (*dto.CPLResponse, error) {
	cpl, err := s.cplRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("CPL tidak ditemukan")
	}

	cpl.Status = req.Status

	if err := s.cplRepo.Update(cpl); err != nil {
		return nil, errors.New("gagal mengupdate status CPL")
	}

	resp := toCPLResponse(cpl)
	return &resp, nil
}

func toCPLResponse(cpl *model.CPL) dto.CPLResponse {
	resp := dto.CPLResponse{
		ID:        cpl.ID,
		Kode:      cpl.Kode,
		Nama:      cpl.Nama,
		Deskripsi: cpl.Deskripsi,
		Status:    cpl.Status,
		Version:   cpl.Version,
		CreatedAt: cpl.CreatedAt,
		UpdatedAt: cpl.UpdatedAt,
		CreatedBy: cpl.CreatedBy,
	}

	if cpl.Creator.ID != "" {
		creator := toUserResponse(&cpl.Creator)
		resp.Creator = &creator
	}

	return resp
}

func (s *cplService) UpdateCPLStatus(id string, status string) (*dto.CPLResponse, error) {
	cpl, err := s.cplRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("CPL tidak ditemukan")
	}

	cpl.Status = status

	if err := s.cplRepo.Update(cpl); err != nil {
		return nil, errors.New("gagal mengupdate status CPL")
	}

	resp := toCPLResponse(cpl)
	return &resp, nil
}

func (s *cplService) GetCPLStatistics() (*dto.CPLStatisticsResponse, error) {
	total, err := s.cplRepo.Count()
	if err != nil {
		return nil, errors.New("gagal mengambil statistik CPL")
	}

	published, _ := s.cplRepo.CountByStatus("published")
	draft, _ := s.cplRepo.CountByStatus("draft")
	archived, _ := s.cplRepo.CountByStatus("archived")

	return &dto.CPLStatisticsResponse{
		TotalCPL:  total,
		Published: published,
		Draft:     draft,
		Archived:  archived,
	}, nil
}

func (s *cplService) GetActiveCPL() ([]dto.CPLResponse, error) {
	cpls, err := s.cplRepo.FindByStatus("published")
	if err != nil {
		return nil, errors.New("gagal mengambil CPL aktif")
	}

	var responses []dto.CPLResponse
	for _, cpl := range cpls {
		responses = append(responses, toCPLResponse(&cpl))
	}

	return responses, nil
}
