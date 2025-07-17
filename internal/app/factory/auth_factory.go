package factory

import (
	auth_handler "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/auth/handler"
	auth_usecase "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/auth/usecase"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
)

func newAuthFactory(userRepo repository.UserRepository) *auth_handler.AuthHandler {
	usecase := auth_usecase.NewAuthUsecase(userRepo)
	handler := auth_handler.NewAuthHandler(usecase)
	return handler
}
