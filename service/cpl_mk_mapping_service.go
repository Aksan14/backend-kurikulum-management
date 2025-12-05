package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"errors"

	"gorm.io/gorm"
)

type CPLMKMappingService interface {
	GetAll(req dto.CPLMKMappingListRequest) (*dto.PaginatedResponse, error)
	GetByID(id string) (*dto.CPLMKMappingResponse, error)
	Upsert(req dto.UpsertCPLMKMappingRequest) (*dto.CPLMKMappingResponse, bool, error)
	Delete(id string) error
}

type cplMKMappingService struct {
	mappingRepo repository.CPLMKMappingRepository
	cplRepo     repository.CPLRepository
	mkRepo      repository.MataKuliahRepository
}

func NewCPLMKMappingService(
	mappingRepo repository.CPLMKMappingRepository,
	cplRepo repository.CPLRepository,
	mkRepo repository.MataKuliahRepository,
) CPLMKMappingService {
	return &cplMKMappingService{
		mappingRepo: mappingRepo,
		cplRepo:     cplRepo,
		mkRepo:      mkRepo,
	}
}

func (s *cplMKMappingService) GetAll(req dto.CPLMKMappingListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Limit > 1000 {
		req.Limit = 1000
	}

	mappings, total, err := s.mappingRepo.FindAll(req.Page, req.Limit, req.CPLID, req.MataKuliahID, req.Level)
	if err != nil {
		return nil, err
	}

	responses := []dto.CPLMKMappingResponse{}
	for _, m := range mappings {
		responses = append(responses, toCPLMKMappingResponse(&m))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *cplMKMappingService) GetByID(id string) (*dto.CPLMKMappingResponse, error) {
	mapping, err := s.mappingRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("mapping tidak ditemukan")
	}

	resp := toCPLMKMappingResponse(mapping)
	return &resp, nil
}

func (s *cplMKMappingService) Upsert(req dto.UpsertCPLMKMappingRequest) (*dto.CPLMKMappingResponse, bool, error) {
	// Validate CPL exists
	_, err := s.cplRepo.FindByID(req.CPLID)
	if err != nil {
		return nil, false, errors.New("CPL tidak ditemukan")
	}

	// Validate Mata Kuliah exists
	_, err = s.mkRepo.FindByID(req.MataKuliahID)
	if err != nil {
		return nil, false, errors.New("mata Kuliah tidak ditemukan")
	}

	// Check if mapping already exists
	existingMapping, err := s.mappingRepo.FindByCPLAndMK(req.CPLID, req.MataKuliahID)

	isNew := false
	var mapping *model.CPLMKMapping

	if err == gorm.ErrRecordNotFound || existingMapping == nil {
		// Create new mapping
		mapping = &model.CPLMKMapping{
			CPLID:        req.CPLID,
			MataKuliahID: req.MataKuliahID,
			Level:        req.Level,
		}
		if err := s.mappingRepo.Create(mapping); err != nil {
			return nil, false, errors.New("gagal membuat mapping")
		}
		isNew = true
	} else {
		// Update existing mapping
		existingMapping.Level = req.Level
		if err := s.mappingRepo.Update(existingMapping); err != nil {
			return nil, false, errors.New("gagal mengupdate mapping")
		}
		mapping = existingMapping
	}

	// Reload with relations
	mapping, _ = s.mappingRepo.FindByID(mapping.ID)
	resp := toCPLMKMappingResponse(mapping)
	return &resp, isNew, nil
}

func (s *cplMKMappingService) Delete(id string) error {
	_, err := s.mappingRepo.FindByID(id)
	if err != nil {
		return errors.New("mapping tidak ditemukan")
	}

	return s.mappingRepo.Delete(id)
}

func toCPLMKMappingResponse(m *model.CPLMKMapping) dto.CPLMKMappingResponse {
	resp := dto.CPLMKMappingResponse{
		ID:           m.ID,
		CPLID:        m.CPLID,
		MataKuliahID: m.MataKuliahID,
		Level:        m.Level,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}

	// Add CPL info if loaded
	if m.CPL.ID != "" {
		resp.CPL = &dto.CPLSimpleResponse{
			ID:   m.CPL.ID,
			Kode: m.CPL.Kode,
			Nama: m.CPL.Nama,
		}
	}

	// Add MataKuliah info if loaded
	if m.MataKuliah.ID != "" {
		resp.MataKuliah = &dto.MKSimpleResponse{
			ID:       m.MataKuliah.ID,
			Kode:     m.MataKuliah.Kode,
			Nama:     m.MataKuliah.Nama,
			Semester: m.MataKuliah.Semester,
		}
	}

	return resp
}
