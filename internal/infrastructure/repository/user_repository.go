package repository

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	FindAll() ([]domain.User, error)
	Create(user *domain.User) error
	Update(id string, user *domain.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(id string, user *domain.User) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}
