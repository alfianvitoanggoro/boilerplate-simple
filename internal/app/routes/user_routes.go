package routes

import (
	user_handler "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/handler"
	"github.com/gofiber/fiber/v2"
)

func NewUserRoutes(routerUser fiber.Router, handler *user_handler.UserHandler) {
	routerUser.Get("/", handler.GetAll)
	routerUser.Get("/:id", handler.GetByID)
	routerUser.Post("/", handler.Create)
	routerUser.Put("/:id", handler.Update)
	routerUser.Delete("/:id", handler.Delete)
}
