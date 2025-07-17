package routes

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/delivery"
	"github.com/gofiber/fiber/v2"
)

func NewUserRoutes(routerUser fiber.Router, handler *delivery.UserHandler) {
	routerUser.Get("/", handler.GetAll)
	routerUser.Get("/:id", handler.GetByID)
	routerUser.Post("/", handler.Create)
	routerUser.Put("/:id", handler.Update)
	routerUser.Delete("/:id", handler.Delete)
}
