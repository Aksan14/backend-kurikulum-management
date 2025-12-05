package controller

import (
	"backend-kurikulum-apps/dto"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileController struct {
	uploadDir string
}

func NewFileController(uploadDir string) *FileController {
	return &FileController{uploadDir: uploadDir}
}

// UploadFile godoc
// @Summary Upload file
// @Description Upload a file
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formance  file true "File to upload"
// @Param type formData string false "File type (image, document, etc.)"
// @Success 200 {object} dto.APIResponse{data=dto.FileUploadResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /files/upload [post]
func (c *FileController) UploadFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "File tidak ditemukan",
			Error:   err.Error(),
		})
		return
	}
	defer file.Close()

	// Validate file size (max 10MB)
	if header.Size > 10*1024*1024 {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Ukuran file maksimal 10MB",
		})
		return
	}

	// Get file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".pdf":  true,
		".doc":  true,
		".docx": true,
		".xls":  true,
		".xlsx": true,
	}

	if !allowedExts[ext] {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Tipe file tidak diizinkan",
		})
		return
	}

	// Generate unique filename
	fileType := ctx.DefaultPostForm("type", "general")
	filename := fmt.Sprintf("%s_%s%s", fileType, uuid.New().String(), ext)

	// Create directory if not exists
	uploadPath := filepath.Join(c.uploadDir, fileType)
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Gagal membuat direktori upload",
		})
		return
	}

	// Create destination file
	destPath := filepath.Join(uploadPath, filename)
	dest, err := os.Create(destPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Gagal menyimpan file",
		})
		return
	}
	defer dest.Close()

	// Copy file content
	if _, err := io.Copy(dest, file); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Gagal menyimpan file",
		})
		return
	}

	// Build file URL
	fileURL := fmt.Sprintf("/uploads/%s/%s", fileType, filename)

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "File berhasil diupload",
		Data: dto.FileUploadResponse{
			StoredName:   filename,
			OriginalName: header.Filename,
			URL:          fileURL,
			FileSize:     header.Size,
			MimeType:     header.Header.Get("Content-Type"),
			FileType:     fileType,
			FilePath:     filepath.Join(uploadPath, filename),
			CreatedAt:    time.Now(),
		},
	})
}

// DeleteFile godoc
// @Summary Delete file
// @Description Delete an uploaded file
// @Tags File
// @Produce json
// @Security BearerAuth
// @Param path query string true "File path to delete"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /files [delete]
func (c *FileController) DeleteFile(ctx *gin.Context) {
	filePath := ctx.Query("path")
	if filePath == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Path file tidak ditemukan",
		})
		return
	}

	// Sanitize path to prevent directory traversal
	cleanPath := filepath.Clean(filePath)
	if strings.Contains(cleanPath, "..") {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Path tidak valid",
		})
		return
	}

	// Build full path
	fullPath := filepath.Join(c.uploadDir, strings.TrimPrefix(cleanPath, "/uploads/"))

	// Check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: "File tidak ditemukan",
		})
		return
	}

	// Delete file
	if err := os.Remove(fullPath); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Gagal menghapus file",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "File berhasil dihapus",
	})
}

// GetFile godoc
// @Summary Get file info
// @Description Get information about uploaded file
// @Tags File
// @Produce json
// @Param path query string true "File path"
// @Success 200 {object} dto.APIResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /files/info [get]
func (c *FileController) GetFileInfo(ctx *gin.Context) {
	filePath := ctx.Query("path")
	if filePath == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Path file tidak ditemukan",
		})
		return
	}

	// Sanitize path
	cleanPath := filepath.Clean(filePath)
	if strings.Contains(cleanPath, "..") {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Path tidak valid",
		})
		return
	}

	// Build full path
	fullPath := filepath.Join(c.uploadDir, strings.TrimPrefix(cleanPath, "/uploads/"))

	// Get file info
	info, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: "File tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan info file",
		Data: map[string]interface{}{
			"name":     info.Name(),
			"size":     info.Size(),
			"modified": info.ModTime(),
			"is_dir":   info.IsDir(),
			"path":     filePath,
		},
	})
}
