package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService service.DocumentService
}

func NewDocumentController(documentService service.DocumentService) *DocumentController {
	return &DocumentController{documentService: documentService}
}

// ========== Template Endpoints ==========

// GetAllTemplates godoc
// @Summary Get all document templates
// @Description Get all document templates dengan pagination
// @Tags Document
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by nama"
// @Param is_active query bool false "Filter by active status"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /documents/templates [get]
func (c *DocumentController) GetAllTemplates(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := dto.DocumentTemplateListRequest{
		Page:      page,
		Limit:     limit,
		Search:    ctx.Query("search"),
		SortBy:    ctx.Query("sort_by"),
		SortOrder: ctx.Query("sort_order"),
	}

	if isActive := ctx.Query("is_active"); isActive != "" {
		val, _ := strconv.ParseBool(isActive)
		req.IsActive = &val
	}

	response, err := c.documentService.GetAllTemplates(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data template dokumen",
		Data:    response,
	})
}

// GetTemplateByID godoc
// @Summary Get template by ID
// @Description Get detail template berdasarkan ID
// @Tags Document
// @Produce json
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} dto.APIResponse{data=dto.DocumentTemplateResponse}
// @Failure 404 {object} dto.ErrorResponse
// @Router /documents/templates/{id} [get]
func (c *DocumentController) GetTemplateByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.documentService.GetTemplateByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data template",
		Data:    response,
	})
}

// CreateTemplate godoc
// @Summary Create document template
// @Description Create new document template (Kaprodi only)
// @Tags Document
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.DocumentTemplateRequest true "Create Template Request"
// @Success 201 {object} dto.APIResponse{data=dto.DocumentTemplateResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /documents/templates [post]
func (c *DocumentController) CreateTemplate(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var req dto.DocumentTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.documentService.CreateTemplate(userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Template berhasil dibuat",
		Data:    response,
	})
}

// UpdateTemplate godoc
// @Summary Update document template
// @Description Update document template (Kaprodi only)
// @Tags Document
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Param request body dto.DocumentTemplateRequest true "Update Template Request"
// @Success 200 {object} dto.APIResponse{data=dto.DocumentTemplateResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /documents/templates/{id} [put]
func (c *DocumentController) UpdateTemplate(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.DocumentTemplateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.documentService.UpdateTemplate(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Template berhasil diupdate",
		Data:    response,
	})
}

// DeleteTemplate godoc
// @Summary Delete document template
// @Description Delete document template (Kaprodi only)
// @Tags Document
// @Produce json
// @Security BearerAuth
// @Param id path string true "Template ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /documents/templates/{id} [delete]
func (c *DocumentController) DeleteTemplate(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.documentService.DeleteTemplate(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Template berhasil dihapus",
	})
}

// ========== Generated Document Endpoints ==========

// GetAllDocuments godoc
// @Summary Get all generated documents
// @Description Get all generated documents dengan pagination
// @Tags Document
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param template_id query string false "Filter by template"
// @Param status query string false "Filter by status"
// @Param tahun query string false "Filter by tahun"
// @Param file_type query string false "Filter by file type"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /documents [get]
func (c *DocumentController) GetAllDocuments(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := dto.GeneratedDocumentListRequest{
		Page:       page,
		Limit:      limit,
		TemplateID: ctx.Query("template_id"),
		Status:     ctx.Query("status"),
		Tahun:      ctx.Query("tahun"),
		FileType:   ctx.Query("file_type"),
		SortBy:     ctx.Query("sort_by"),
		SortOrder:  ctx.Query("sort_order"),
	}

	response, err := c.documentService.GetAllDocuments(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data dokumen",
		Data:    response,
	})
}

// GetDocumentByID godoc
// @Summary Get document by ID
// @Description Get detail generated document berdasarkan ID
// @Tags Document
// @Produce json
// @Security BearerAuth
// @Param id path string true "Document ID"
// @Success 200 {object} dto.APIResponse{data=dto.GeneratedDocumentResponse}
// @Failure 404 {object} dto.ErrorResponse
// @Router /documents/{id} [get]
func (c *DocumentController) GetDocumentByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.documentService.GetDocumentByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data dokumen",
		Data:    response,
	})
}

// GenerateDocument godoc
// @Summary Generate document
// @Description Generate new document from template
// @Tags Document
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.GenerateDocumentRequest true "Generate Document Request"
// @Success 201 {object} dto.APIResponse{data=dto.GeneratedDocumentResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /documents/generate [post]
func (c *DocumentController) GenerateDocument(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var req dto.GenerateDocumentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.documentService.GenerateDocument(userID.(string), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "Dokumen sedang diproses",
		Data:    response,
	})
}

// DeleteDocument godoc
// @Summary Delete document
// @Description Delete generated document
// @Tags Document
// @Produce json
// @Security BearerAuth
// @Param id path string true "Document ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /documents/{id} [delete]
func (c *DocumentController) DeleteDocument(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.documentService.DeleteDocument(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Dokumen berhasil dihapus",
	})
}
