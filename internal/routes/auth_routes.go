package routes

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/delivery"
	"github.com/gofiber/fiber/v2"
)

func NewAuthRoutes(routerAuth fiber.Router, handler *delivery.AuthHandler) {
	routerAuth.Post("/register", handler.Register)
	routerAuth.Post("/login", handler.Login)
}
