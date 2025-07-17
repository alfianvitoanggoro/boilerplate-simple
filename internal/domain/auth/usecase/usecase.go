package auth_usecase

import (
	"errors"
	"os"
	"time"

	user_model "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/model"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/hash"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	Register(username, password, role string) error
	Login(username, password string) (string, error)
}

type authUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(r repository.UserRepository) AuthUsecase {
	return &authUsecase{repo: r}
}

func (u *authUsecase) Register(username, password, role string) error {
	hashPwd, err := hash.HashPassword(password)
	if err != nil {
		return err
	}
	user := &user_model.User{
		Username: username,
		Password: hashPwd,
		Role:     user_model.Role(role),
	}
	return u.repo.Create(user)
}

func (u *authUsecase) Login(username, password string) (string, error) {
	user, err := u.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !hash.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
