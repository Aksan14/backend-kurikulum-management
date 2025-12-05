package service

import (
	"backend-kurikulum-apps/config"
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/helper"
	"backend-kurikulum-apps/model"
	"backend-kurikulum-apps/repository"
	"errors"
	"fmt"
	"time"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
	Register(req dto.RegisterRequest) (*dto.UserResponse, error)
	RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	Logout(userID string) error
	ChangePassword(userID string, req dto.ChangePasswordRequest) error
	GetProfile(userID string) (*dto.UserResponse, error)
	UpdateProfile(userID string, req dto.UpdateProfileRequest) (*dto.UserResponse, error)
}

type authService struct {
	userRepo    repository.UserRepository
	refreshRepo repository.RefreshTokenRepository
}

func NewAuthService(userRepo repository.UserRepository, refreshRepo repository.RefreshTokenRepository) AuthService {
	return &authService{
		userRepo:    userRepo,
		refreshRepo: refreshRepo,
	}
}

func (s *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	if user.Status != "active" {
		return nil, errors.New("akun tidak aktif")
	}

	if !helper.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("email atau password salah")
	}

	// Generate access token
	accessToken, err := helper.GenerateJWT(
		user.ID,
		user.Email,
		user.Role,
		config.JWT.Secret,
		config.JWT.AccessExpires,
	)
	if err != nil {
		return nil, errors.New("gagal membuat token")
	}

	// Generate refresh token
	refreshToken, refreshHash, err := helper.GenerateRefreshToken()
	if err != nil {
		return nil, errors.New("gagal membuat refresh token")
	}

	// Revoke old refresh tokens
	s.refreshRepo.RevokeByUserID(user.ID)

	// Save new refresh token
	rt := &model.RefreshToken{
		UserID:    user.ID,
		Token:     refreshHash,
		ExpiresAt: time.Now().Add(config.JWT.RefreshExpires),
	}
	if err := s.refreshRepo.Create(rt); err != nil {
		return nil, errors.New("gagal menyimpan refresh token")
	}

	// Update last login
	s.userRepo.UpdateLastLogin(user.ID)

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(config.JWT.AccessExpires.Seconds()),
		TokenType:    "Bearer",
		User:         toUserResponse(user),
	}, nil
}

func (s *authService) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	// Check if email already exists
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	// Hash password
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi password")
	}

	user := &model.User{
		Nama:         req.Nama,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Role:         req.Role,
		NIP:          req.NIP,
		Phone:        req.Phone,
		Status:       "active",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("gagal membuat user: %v", err)
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *authService) RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	tokenHash := helper.HashToken(req.RefreshToken)

	rt, err := s.refreshRepo.FindByToken(tokenHash)
	if err != nil {
		return nil, errors.New("refresh token tidak valid")
	}

	if rt.Revoked || rt.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token sudah expired")
	}

	user, err := s.userRepo.FindByID(rt.UserID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// Generate new access token
	accessToken, err := helper.GenerateJWT(
		user.ID,
		user.Email,
		user.Role,
		config.JWT.Secret,
		config.JWT.AccessExpires,
	)
	if err != nil {
		return nil, errors.New("gagal membuat token")
	}

	// Generate new refresh token
	newRefreshToken, newRefreshHash, err := helper.GenerateRefreshToken()
	if err != nil {
		return nil, errors.New("gagal membuat refresh token")
	}

	// Revoke old token
	s.refreshRepo.RevokeByID(rt.ID)

	// Save new refresh token
	newRT := &model.RefreshToken{
		UserID:    user.ID,
		Token:     newRefreshHash,
		ExpiresAt: time.Now().Add(config.JWT.RefreshExpires),
	}
	if err := s.refreshRepo.Create(newRT); err != nil {
		return nil, errors.New("gagal menyimpan refresh token")
	}

	return &dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    int64(config.JWT.AccessExpires.Seconds()),
		TokenType:    "Bearer",
	}, nil
}

func (s *authService) Logout(userID string) error {
	return s.refreshRepo.RevokeByUserID(userID)
}

func (s *authService) ChangePassword(userID string, req dto.ChangePasswordRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	if !helper.CheckPassword(req.OldPassword, user.PasswordHash) {
		return errors.New("password lama salah")
	}

	hashedPassword, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("gagal mengenkripsi password")
	}

	user.PasswordHash = hashedPassword
	return s.userRepo.Update(user)
}

func (s *authService) GetProfile(userID string) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func (s *authService) UpdateProfile(userID string, req dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.AvatarURL != nil {
		user.AvatarURL = req.AvatarURL
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, errors.New("gagal mengupdate profil")
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func toUserResponse(user *model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Nama:      user.Nama,
		Email:     user.Email,
		NIP:       user.NIP,
		Role:      user.Role,
		Status:    user.Status,
		Phone:     user.Phone,
		AvatarURL: user.AvatarURL,
		LastLogin: user.LastLogin,
		CreatedAt: user.CreatedAt,
	}
}
