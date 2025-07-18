package repository

import (
	user_model "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*user_model.User, error)
	FindByID(id string) (*user_model.User, error)
	FindAllPaginated(page, limit int) ([]user_model.User, int, error)
	Create(user *user_model.User) error
	Update(id string, user *user_model.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(username string) (*user_model.User, error) {
	var user user_model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id string) (*user_model.User, error) {
	var user user_model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAllPaginated(page, limit int) ([]user_model.User, int, error) {
	var users []user_model.User
	var total int64
	r.db.Model(&user_model.User{}).Count(&total)
	offset := (page - 1) * limit
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, int(total), nil
}

func (r *userRepository) Update(id string, user *user_model.User) error {
	return r.db.Model(&user_model.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&user_model.User{}, id).Error
}

func (r *userRepository) Create(user *user_model.User) error {
	return r.db.Create(user).Error
}
