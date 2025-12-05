package controller

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users dengan pagination
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by nama or email"
// @Param role query string false "Filter by role"
// @Param is_active query bool false "Filter by active status"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} dto.APIResponse{data=dto.PaginatedResponse}
// @Failure 401 {object} dto.ErrorResponse
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	req := dto.UserListRequest{
		Page:      page,
		Limit:     limit,
		Search:    ctx.Query("search"),
		Role:      ctx.Query("role"),
		Status:    ctx.Query("status"),
		SortBy:    ctx.Query("sort_by"),
		SortOrder: ctx.Query("sort_order"),
	}

	response, err := c.userService.GetAll(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data users",
		Data:    response,
	})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get detail user berdasarkan ID
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} dto.APIResponse{data=dto.UserResponse}
// @Failure 404 {object} dto.ErrorResponse
// @Router /users/{id} [get]
func (c *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.userService.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data user",
		Data:    response,
	})
}

// CreateUser godoc
// @Summary Create new user
// @Description Create new user (Admin/Kaprodi only)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateUserRequest true "Create User Request"
// @Success 201 {object} dto.APIResponse{data=dto.UserResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.userService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.APIResponse{
		Success: true,
		Message: "User berhasil dibuat",
		Data:    response,
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update data user (Admin/Kaprodi only)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body dto.UpdateUserRequest true "Update User Request"
// @Success 200 {object} dto.APIResponse{data=dto.UserResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Data tidak valid",
			Error:   err.Error(),
		})
		return
	}

	response, err := c.userService.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "User berhasil diupdate",
		Data:    response,
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user (Admin/Kaprodi only)
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} dto.APIResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.userService.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "User berhasil dihapus",
	})
}

// GetDosen godoc
// @Summary Get all dosen
// @Description Get all users with role dosen
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.APIResponse{data=[]dto.UserResponse}
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/dosen [get]
func (c *UserController) GetDosen(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "100"))

	response, err := c.userService.GetDosen(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Berhasil mendapatkan data dosen",
		Data:    response,
	})
}

// ToggleUserStatus godoc
// @Summary Toggle user active status
// @Description Activate or deactivate user
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} dto.APIResponse{data=dto.UserResponse}
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id}/toggle-status [patch]
func (c *UserController) ToggleUserStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.userService.ToggleUserStatus(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.APIResponse{
		Success: true,
		Message: "Status user berhasil diubah",
		Data:    response,
	})
}
