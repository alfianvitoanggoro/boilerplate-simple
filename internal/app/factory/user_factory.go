package factory

import (
	user_handler "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/handler"
	user_usecase "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/usecase"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
)

func newUserFactory(userRepo repository.UserRepository) *user_handler.UserHandler {
	usecase := user_usecase.NewUserUsecase(userRepo)
	handler := user_handler.NewUserHandler(usecase)
	return handler
}
