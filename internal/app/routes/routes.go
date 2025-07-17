package routes

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/app/factory"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(app *fiber.App, container *factory.Container) {
	api := app.Group("/api")

	// Register auth routes
	auth := api.Group("/auth")
	NewAuthRoutes(auth, container.AuthHandler)

	// Protected users endpoints (JWT + RBAC admin)
	users := api.Group("/users", middleware.JWTMiddleware(), middleware.RBACMiddleware("admin", "teacher"))
	NewUserRoutes(users, container.UserHandler)
}
