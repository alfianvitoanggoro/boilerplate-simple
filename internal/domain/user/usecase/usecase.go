package user_usecase

import (
	user_dto "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/dto"
	user_model "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/model"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
)

type UserUsecase interface {
	GetAll() ([]user_dto.UserResponse, error)
	GetByID(id string) (user_dto.UserResponse, error)
	Create(req user_dto.UserCreateRequest) (user_dto.UserResponse, error)
	Update(id string, req user_dto.UserUpdateRequest) (user_dto.UserResponse, error)
	Delete(id string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) GetAll() ([]user_dto.UserResponse, error) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var resp []user_dto.UserResponse
	for _, user := range users {
		resp = append(resp, user_dto.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     string(user.Role),
		})
	}
	return resp, nil
}

func (u *userUsecase) GetByID(id string) (user_dto.UserResponse, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return user_dto.UserResponse{}, err
	}
	return user_dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     string(user.Role),
	}, nil
}

func (u *userUsecase) Create(req user_dto.UserCreateRequest) (user_dto.UserResponse, error) {
	user := &user_model.User{
		Username: req.Username,
		Password: req.Password, // You may want to hash this!
		Role:     user_model.Role(req.Role),
	}
	if err := u.repo.Create(user); err != nil {
		return user_dto.UserResponse{}, err
	}
	return user_dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     string(user.Role),
	}, nil
}

func (u *userUsecase) Update(id string, req user_dto.UserUpdateRequest) (user_dto.UserResponse, error) {
	user := &user_model.User{
		Username: req.Username,
		Role:     user_model.Role(req.Role),
	}
	if err := u.repo.Update(id, user); err != nil {
		return user_dto.UserResponse{}, err
	}
	updated, err := u.repo.FindByID(id)
	if err != nil {
		return user_dto.UserResponse{}, err
	}

	return user_dto.UserResponse{
		ID:       updated.ID,
		Username: updated.Username,
		Role:     string(updated.Role),
	}, nil
}

func (u *userUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
