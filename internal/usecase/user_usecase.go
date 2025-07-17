package usecase

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/domain"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/dto"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
)

type UserUsecase interface {
	GetAll() ([]dto.UserResponse, error)
	GetByID(id string) (dto.UserResponse, error)
	Create(req dto.UserCreateRequest) (dto.UserResponse, error)
	Update(id string, req dto.UserUpdateRequest) (dto.UserResponse, error)
	Delete(id string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetAll() ([]dto.UserResponse, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var resp []dto.UserResponse
	for _, user := range users {
		resp = append(resp, dto.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     string(user.Role),
		})
	}
	return resp, nil
}

func (u *userUsecase) GetByID(id string) (dto.UserResponse, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     string(user.Role),
	}, nil
}

func (u *userUsecase) Create(req dto.UserCreateRequest) (dto.UserResponse, error) {
	user := &domain.User{
		Username: req.Username,
		Password: req.Password, // You may want to hash this!
		Role:     domain.Role(req.Role),
	}
	if err := u.repo.Create(user); err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     string(user.Role),
	}, nil
}

func (u *userUsecase) Update(id string, req dto.UserUpdateRequest) (dto.UserResponse, error) {
	user := &domain.User{
		Username: req.Username,
		Role:     domain.Role(req.Role),
	}
	if err := u.repo.Update(id, user); err != nil {
		return dto.UserResponse{}, err
	}
	updated, err := u.repo.FindByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:       updated.ID,
		Username: updated.Username,
		Role:     string(updated.Role),
	}, nil
}

func (u *userUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
