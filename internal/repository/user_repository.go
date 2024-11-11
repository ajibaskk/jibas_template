package repository

import (
	"jibas-template/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

// Create implements domain.UserRepository.
func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// GetAll implements domain.UserRepository.
func (r *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}
