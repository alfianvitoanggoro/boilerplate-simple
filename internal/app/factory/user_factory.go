package factory

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/delivery"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/usecase"
)

func newUserFactory(userRepo repository.UserRepository) *delivery.UserHandler {
	usecase := usecase.NewUserUsecase(userRepo)
	handler := delivery.NewUserHandler(usecase)
	return handler
}
