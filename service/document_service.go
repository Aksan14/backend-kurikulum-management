package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"encoding/json"
	"errors"
)

type DocumentService interface {
	// Templates
	GetAllTemplates(req dto.DocumentTemplateListRequest) (*dto.PaginatedResponse, error)
	GetTemplateByID(id string) (*dto.DocumentTemplateResponse, error)
	CreateTemplate(userID string, req dto.DocumentTemplateRequest) (*dto.DocumentTemplateResponse, error)
	UpdateTemplate(id string, req dto.DocumentTemplateRequest) (*dto.DocumentTemplateResponse, error)
	DeleteTemplate(id string) error

	// Generated Documents
	GetAllDocuments(req dto.GeneratedDocumentListRequest) (*dto.PaginatedResponse, error)
	GetDocumentByID(id string) (*dto.GeneratedDocumentResponse, error)
	GenerateDocument(userID string, req dto.GenerateDocumentRequest) (*dto.GeneratedDocumentResponse, error)
	DeleteDocument(id string) error
}

type documentService struct {
	templateRepo repository.DocumentTemplateRepository
	docRepo      repository.GeneratedDocumentRepository
}

func NewDocumentService(
	templateRepo repository.DocumentTemplateRepository,
	docRepo repository.GeneratedDocumentRepository,
) DocumentService {
	return &documentService{
		templateRepo: templateRepo,
		docRepo:      docRepo,
	}
}

// Template methods
func (s *documentService) GetAllTemplates(req dto.DocumentTemplateListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	templates, total, err := s.templateRepo.FindAll(
		req.Page, req.Limit, req.Search, req.IsActive, req.SortBy, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.DocumentTemplateResponse
	for _, t := range templates {
		responses = append(responses, toDocumentTemplateResponse(&t))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *documentService) GetTemplateByID(id string) (*dto.DocumentTemplateResponse, error) {
	template, err := s.templateRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("template tidak ditemukan")
	}

	resp := toDocumentTemplateResponse(template)
	return &resp, nil
}

func (s *documentService) CreateTemplate(userID string, req dto.DocumentTemplateRequest) (*dto.DocumentTemplateResponse, error) {
	sectionsJSON, _ := json.Marshal(req.Sections)

	version := req.Version
	if version == "" {
		version = "1.0"
	}

	template := &model.DocumentTemplate{
		Nama:      req.Nama,
		Deskripsi: req.Deskripsi,
		Sections:  sectionsJSON,
		FileURL:   req.FileURL,
		Version:   version,
		IsActive:  req.IsActive,
		CreatedBy: userID,
	}

	if err := s.templateRepo.Create(template); err != nil {
		return nil, errors.New("gagal membuat template")
	}

	resp := toDocumentTemplateResponse(template)
	return &resp, nil
}

func (s *documentService) UpdateTemplate(id string, req dto.DocumentTemplateRequest) (*dto.DocumentTemplateResponse, error) {
	template, err := s.templateRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("template tidak ditemukan")
	}

	sectionsJSON, _ := json.Marshal(req.Sections)

	template.Nama = req.Nama
	template.Deskripsi = req.Deskripsi
	template.Sections = sectionsJSON
	template.FileURL = req.FileURL
	if req.Version != "" {
		template.Version = req.Version
	}
	template.IsActive = req.IsActive

	if err := s.templateRepo.Update(template); err != nil {
		return nil, errors.New("gagal mengupdate template")
	}

	resp := toDocumentTemplateResponse(template)
	return &resp, nil
}

func (s *documentService) DeleteTemplate(id string) error {
	_, err := s.templateRepo.FindByID(id)
	if err != nil {
		return errors.New("template tidak ditemukan")
	}

	return s.templateRepo.Delete(id)
}

// Generated Document methods
func (s *documentService) GetAllDocuments(req dto.GeneratedDocumentListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	docs, total, err := s.docRepo.FindAll(
		req.Page, req.Limit, req.TemplateID, req.Status, req.Tahun, req.FileType, req.SortBy, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.GeneratedDocumentResponse
	for _, d := range docs {
		responses = append(responses, toGeneratedDocumentResponse(&d))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *documentService) GetDocumentByID(id string) (*dto.GeneratedDocumentResponse, error) {
	doc, err := s.docRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("dokumen tidak ditemukan")
	}

	resp := toGeneratedDocumentResponse(doc)
	return &resp, nil
}

func (s *documentService) GenerateDocument(userID string, req dto.GenerateDocumentRequest) (*dto.GeneratedDocumentResponse, error) {
	// Get template
	template, err := s.templateRepo.FindByID(req.TemplateID)
	if err != nil {
		return nil, errors.New("template tidak ditemukan")
	}

	sectionsJSON, _ := json.Marshal(req.Sections)
	generationDataJSON, _ := json.Marshal(req.GenerationData)

	doc := &model.GeneratedDocument{
		TemplateID:     req.TemplateID,
		TemplateName:   template.Nama,
		Tahun:          req.Tahun,
		Status:         "processing",
		FileType:       req.FileType,
		Sections:       sectionsJSON,
		GenerationData: generationDataJSON,
		Progress:       0,
		CreatedBy:      userID,
	}

	if err := s.docRepo.Create(doc); err != nil {
		return nil, errors.New("gagal membuat dokumen")
	}

	// TODO: Implement actual document generation in background
	// For now, just mark it as ready (in real implementation, this would be async)
	go s.processDocumentGeneration(doc.ID)

	doc, _ = s.docRepo.FindByID(doc.ID)
	resp := toGeneratedDocumentResponse(doc)
	return &resp, nil
}

func (s *documentService) processDocumentGeneration(docID string) {
	// Simulate document generation
	// In real implementation, this would generate the actual document

	// Update progress
	s.docRepo.UpdateProgress(docID, 50)

	// Generate file URL (placeholder)
	fileURL := "/uploads/documents/" + docID + ".pdf"
	var fileSize int64 = 1024 * 100 // 100KB placeholder

	s.docRepo.UpdateStatus(docID, "ready", &fileURL, &fileSize, nil)
}

func (s *documentService) DeleteDocument(id string) error {
	_, err := s.docRepo.FindByID(id)
	if err != nil {
		return errors.New("dokumen tidak ditemukan")
	}

	return s.docRepo.Delete(id)
}

// Response converters
func toDocumentTemplateResponse(t *model.DocumentTemplate) dto.DocumentTemplateResponse {
	var sections []string
	json.Unmarshal(t.Sections, &sections)

	resp := dto.DocumentTemplateResponse{
		ID:        t.ID,
		Nama:      t.Nama,
		Deskripsi: t.Deskripsi,
		Sections:  sections,
		FileURL:   t.FileURL,
		Version:   t.Version,
		IsActive:  t.IsActive,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		CreatedBy: t.CreatedBy,
	}

	if t.Creator.ID != "" {
		creator := toUserResponse(&t.Creator)
		resp.Creator = &creator
	}

	return resp
}

func toGeneratedDocumentResponse(d *model.GeneratedDocument) dto.GeneratedDocumentResponse {
	var sections []string
	var generationData map[string]interface{}
	json.Unmarshal(d.Sections, &sections)
	json.Unmarshal(d.GenerationData, &generationData)

	resp := dto.GeneratedDocumentResponse{
		ID:             d.ID,
		TemplateID:     d.TemplateID,
		TemplateName:   d.TemplateName,
		Tahun:          d.Tahun,
		Status:         d.Status,
		FileURL:        d.FileURL,
		FileType:       d.FileType,
		FileSize:       d.FileSize,
		Sections:       sections,
		GenerationData: generationData,
		Progress:       d.Progress,
		ErrorMessage:   d.ErrorMessage,
		CreatedAt:      d.CreatedAt,
		CompletedAt:    d.CompletedAt,
		CreatedBy:      d.CreatedBy,
	}

	if d.Template.ID != "" {
		template := toDocumentTemplateResponse(&d.Template)
		resp.Template = &template
	}

	return resp
}
