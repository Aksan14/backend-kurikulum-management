package repository

import (
	"backend-kurikulum-apps/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindAll(page, limit int, search, role, status, sortBy, sortOrder string) ([]model.User, int64, error)
	Update(user *model.User) error
	Delete(id string) error
	UpdateLastLogin(id string) error
	CountByRole(role string) (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(page, limit int, search, role, status, sortBy, sortOrder string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.Model(&model.User{})

	if search != "" {
		query = query.Where("nama LIKE ? OR email LIKE ? OR nip LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	if role != "" {
		query = query.Where("role = ?", role)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	offset := (page - 1) * limit
	err := query.Order(sortBy + " " + sortOrder).Offset(offset).Limit(limit).Find(&users).Error
	return users, total, err
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&model.User{}, "id = ?", id).Error
}

func (r *userRepository) UpdateLastLogin(id string) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("last_login", gorm.Expr("NOW()")).Error
}

func (r *userRepository) CountByRole(role string) (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("role = ? AND status = ?", role, "active").Count(&count).Error
	return count, err
}
