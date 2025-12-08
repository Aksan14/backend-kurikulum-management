package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type CPLAssignmentService interface {
	GetAllAssignments(req dto.CPLAssignmentListRequest) (*dto.PaginatedResponse, error)
	GetAssignmentByID(id string) (*dto.CPLAssignmentResponse, error)
	GetByDosenID(dosenID string, page, limit int, status string) (*dto.PaginatedResponse, error)
	GetAssignmentsByDosen(dosenID string, status string) ([]dto.CPLAssignmentResponse, error)
	GetAssignmentsByCPL(cplID string) ([]dto.CPLAssignmentResponse, error)
	CreateAssignment(createdBy string, req dto.CreateCPLAssignmentRequest) ([]dto.CPLAssignmentResponse, error)
	UpdateAssignmentStatus(id string, req dto.UpdateCPLAssignmentStatusRequest) (*dto.CPLAssignmentResponse, error)
	DeleteAssignment(id string) error
}

type cplAssignmentService struct {
	assignmentRepo repository.CPLAssignmentRepository
	cplRepo        repository.CPLRepository
	mappingRepo    repository.CPLMKMappingRepository
	notifService   NotificationService
}

func NewCPLAssignmentService(
	assignmentRepo repository.CPLAssignmentRepository,
	cplRepo repository.CPLRepository,
	mappingRepo repository.CPLMKMappingRepository,
	notifService NotificationService,
) CPLAssignmentService {
	return &cplAssignmentService{
		assignmentRepo: assignmentRepo,
		cplRepo:        cplRepo,
		mappingRepo:    mappingRepo,
		notifService:   notifService,
	}
}

func (s *cplAssignmentService) GetAllAssignments(req dto.CPLAssignmentListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	assignments, total, err := s.assignmentRepo.FindAll(
		req.Page, req.Limit, req.CPLID, req.DosenID, req.Status, req.SortBy, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.CPLAssignmentResponse
	for _, a := range assignments {
		responses = append(responses, s.toCPLAssignmentResponse(&a))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *cplAssignmentService) GetAssignmentByID(id string) (*dto.CPLAssignmentResponse, error) {
	assignment, err := s.assignmentRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("assignment tidak ditemukan")
	}

	resp := s.toCPLAssignmentResponse(assignment)
	return &resp, nil
}

func (s *cplAssignmentService) GetByDosenID(dosenID string, page, limit int, status string) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	assignments, total, err := s.assignmentRepo.FindByDosenID(dosenID, page, limit, status)
	if err != nil {
		return nil, err
	}

	var responses []dto.CPLAssignmentResponse
	for _, a := range assignments {
		responses = append(responses, s.toCPLAssignmentResponse(&a))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: (total + int64(limit) - 1) / int64(limit),
	}, nil
}

func (s *cplAssignmentService) CreateAssignment(createdBy string, req dto.CreateCPLAssignmentRequest) ([]dto.CPLAssignmentResponse, error) {
	var responses []dto.CPLAssignmentResponse
	var cplIDs []string

	// If no CPL IDs provided, get them from mata kuliah mapping
	if len(req.CPLIDs) == 0 && req.MataKuliahID != nil {
		mappings, err := s.mappingRepo.FindByMataKuliahID(*req.MataKuliahID)
		if err != nil {
			return nil, errors.New("gagal mendapatkan mapping CPL untuk mata kuliah")
		}
		if len(mappings) == 0 {
			return nil, errors.New("tidak ada CPL yang ter-mapping dengan mata kuliah ini")
		}
		for _, mapping := range mappings {
			cplIDs = append(cplIDs, mapping.CPLID)
		}
	} else {
		cplIDs = req.CPLIDs
	}

	// Check for duplicate assignments for each CPL ID
	if req.MataKuliahID != nil {
		for _, cplID := range cplIDs {
			isDuplicate, _ := s.assignmentRepo.CheckDuplicateAssignment(cplID, req.DosenID, *req.MataKuliahID)
			if isDuplicate {
				return nil, errors.New("penugasan sudah ada untuk CPL " + cplID)
			}
		}
	}

	// Convert CPL IDs to JSON
	cplIDsJSON, err := json.Marshal(cplIDs)
	if err != nil {
		return nil, errors.New("gagal mengkonversi CPL IDs ke JSON")
	}

	// Create single assignment with all CPL IDs
	assignment := &model.CPLAssignment{
		CPLIDs:       model.JSON(cplIDsJSON),
		DosenID:      req.DosenID,
		MataKuliah:   req.MataKuliah,
		MataKuliahID: req.MataKuliahID,
		Status:       "assigned",
		AssignedAt:   time.Now(),
		Deadline:     req.Deadline,
		Catatan:      req.Catatan,
		AssignedBy:   createdBy,
	}

	if err := s.assignmentRepo.Create(assignment); err != nil {
		return nil, errors.New("gagal membuat assignment")
	}

	// Reload with relations
	assignment, _ = s.assignmentRepo.FindByID(assignment.ID)
	resp := s.toCPLAssignmentResponse(assignment)
	responses = append(responses, resp)

	// Send notification to dosen
	mataKuliahName := ""
	if req.MataKuliah != nil {
		mataKuliahName = *req.MataKuliah
	}
	if s.notifService != nil && len(responses) > 0 {
		s.notifService.Create(dto.CreateNotificationRequest{
			UserID:      req.DosenID,
			Title:       "Penugasan CPL Baru",
			Message:     fmt.Sprintf("Anda mendapat penugasan %d CPL baru untuk mata kuliah %s", len(cplIDs), mataKuliahName),
			Type:        "assignment",
			RelatedID:   &responses[0].ID,
			RelatedType: stringPtr("cpl_assignment"),
			Priority:    "high",
		})
	}

	return responses, nil
}

func (s *cplAssignmentService) UpdateAssignmentStatus(id string, req dto.UpdateCPLAssignmentStatusRequest) (*dto.CPLAssignmentResponse, error) {
	assignment, err := s.assignmentRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("assignment tidak ditemukan")
	}

	now := time.Now()
	assignment.Status = req.Status

	switch req.Status {
	case "accepted":
		assignment.ResponseAt = &now
	case "completed":
		assignment.CompletedAt = &now
	case "rejected":
		assignment.ResponseAt = &now
		assignment.RejectionReason = req.RejectionReason
	case "cancelled":
		assignment.ResponseAt = &now
		assignment.RejectionReason = req.RejectionReason
	}

	if req.Catatan != nil {
		assignment.Catatan = req.Catatan
	}

	if err := s.assignmentRepo.Update(assignment); err != nil {
		return nil, errors.New("gagal mengupdate status")
	}

	resp := s.toCPLAssignmentResponse(assignment)
	return &resp, nil
}

func (s *cplAssignmentService) DeleteAssignment(id string) error {
	_, err := s.assignmentRepo.FindByID(id)
	if err != nil {
		return errors.New("assignment tidak ditemukan")
	}

	return s.assignmentRepo.Delete(id)
}

func (s *cplAssignmentService) GetAssignmentsByDosen(dosenID string, status string) ([]dto.CPLAssignmentResponse, error) {
	assignments, _, err := s.assignmentRepo.FindByDosenID(dosenID, 1, 1000, status)
	if err != nil {
		return nil, err
	}

	var responses []dto.CPLAssignmentResponse
	for _, a := range assignments {
		responses = append(responses, s.toCPLAssignmentResponse(&a))
	}

	return responses, nil
}

func (s *cplAssignmentService) GetAssignmentsByCPL(cplID string) ([]dto.CPLAssignmentResponse, error) {
	assignments, _, err := s.assignmentRepo.FindAll(1, 1000, cplID, "", "", "", "")
	if err != nil {
		return nil, err
	}

	var responses []dto.CPLAssignmentResponse
	for _, a := range assignments {
		responses = append(responses, s.toCPLAssignmentResponse(&a))
	}

	return responses, nil
}

func (s *cplAssignmentService) toCPLAssignmentResponse(a *model.CPLAssignment) dto.CPLAssignmentResponse {
	var cplIDs []string
	if len(a.CPLIDs) > 0 {
		json.Unmarshal(a.CPLIDs, &cplIDs)
	}

	resp := dto.CPLAssignmentResponse{
		ID:              a.ID,
		CPLIDs:          cplIDs,
		DosenID:         a.DosenID,
		MataKuliah:      a.MataKuliah,
		MataKuliahID:    a.MataKuliahID,
		Deadline:        a.Deadline,
		Status:          a.Status,
		Catatan:         a.Catatan,
		RejectionReason: a.RejectionReason,
		AssignedBy:      a.AssignedBy,
		AssignedAt:      a.AssignedAt,
		ResponseAt:      a.ResponseAt,
		CompletedAt:     a.CompletedAt,
	}

	if len(cplIDs) > 0 {
		// Fetch CPLs by IDs
		cpls, err := s.cplRepo.FindByIDs(cplIDs)
		if err == nil {
			for _, cpl := range cpls {
				resp.CPLs = append(resp.CPLs, toCPLResponse(&cpl))
			}
		}
	}

	if a.Dosen.ID != "" {
		dosen := toUserResponse(&a.Dosen)
		resp.Dosen = &dosen
	}

	if a.MataKuliahRef.ID != "" {
		mk := toMataKuliahResponse(&a.MataKuliahRef)
		resp.MataKuliahRef = &mk
	}

	if a.Assigner.ID != "" {
		assigner := toUserResponse(&a.Assigner)
		resp.Assigner = &assigner
	}

	return resp
}

func stringPtr(s string) *string {
	return &s
}
