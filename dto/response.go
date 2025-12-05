package dto

// PaginatedResponse generic pagination response
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalItems int64       `json:"total_items"`
	TotalPages int64       `json:"total_pages"`
}

// APIResponse standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// ErrorResponse error response
type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

// ValidationError validation error detail
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewSuccessResponse creates a success response
func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates an error response
func NewErrorResponse(message string, err interface{}) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}

// NewPaginatedResponse creates a paginated response
func NewPaginatedResponse(data interface{}, page, limit int, totalItems int64) PaginatedResponse {
	totalPages := totalItems / int64(limit)
	if totalItems%int64(limit) > 0 {
		totalPages++
	}

	return PaginatedResponse{
		Data:       data,
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}
