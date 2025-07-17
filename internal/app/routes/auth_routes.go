package routes

import (
	auth_handler "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/auth/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAuthRoutes(routerAuth fiber.Router, handler *auth_handler.AuthHandler) {
	routerAuth.Post("/register", handler.Register)
	routerAuth.Post("/login", handler.Login)
}
