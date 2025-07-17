package factory

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/delivery"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/usecase"
)

func newAuthFactory(userRepo repository.UserRepository) *delivery.AuthHandler {
	usecase := usecase.NewAuthUsecase(userRepo)
	handler := delivery.NewAuthHandler(usecase)
	return handler
}
