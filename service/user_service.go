package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/helper"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"errors"
)

type UserService interface {
	GetAll(req dto.UserListRequest) (*dto.PaginatedResponse, error)
	GetByID(id string) (*dto.UserResponse, error)
	Create(req dto.CreateUserRequest) (*dto.UserResponse, error)
	Update(id string, req dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(id string) error
	GetDosen(page, limit int) (*dto.PaginatedResponse, error)
	ToggleUserStatus(id string) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAll(req dto.UserListRequest) (*dto.PaginatedResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	users, total, err := s.userRepo.FindAll(
		req.Page, req.Limit, req.Search, req.Role, req.Status, req.SortBy, req.SortOrder,
	)
	if err != nil {
		return nil, err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, toUserResponse(&user))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalItems: total,
		TotalPages: (total + int64(req.Limit) - 1) / int64(req.Limit),
	}, nil
}

func (s *userService) GetByID(id string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *userService) Create(req dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Check if email exists
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi password")
	}

	user := &model.User{
		Nama:         req.Nama,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Role:         req.Role,
		Status:       "active",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("gagal membuat user")
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *userService) Update(id string, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.Email != "" && req.Email != user.Email {
		existing, _ := s.userRepo.FindByEmail(req.Email)
		if existing != nil {
			return nil, errors.New("email sudah digunakan")
		}
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if req.NIP != nil {
		user.NIP = req.NIP
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.Status != "" {
		user.Status = req.Status
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("gagal mengupdate user")
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *userService) Delete(id string) error {
	_, err := s.userRepo.FindByID(id)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	return s.userRepo.Delete(id)
}

func (s *userService) GetDosen(page, limit int) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 100
	}

	users, total, err := s.userRepo.FindAll(page, limit, "", "dosen", "active", "nama", "asc")
	if err != nil {
		return nil, err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, toUserResponse(&user))
	}

	return &dto.PaginatedResponse{
		Data:       responses,
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: (total + int64(limit) - 1) / int64(limit),
	}, nil
}

func (s *userService) ToggleUserStatus(id string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if user.Status == "active" {
		user.Status = "inactive"
	} else {
		user.Status = "active"
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("gagal mengubah status user")
	}

	resp := toUserResponse(user)
	return &resp, nil
}
