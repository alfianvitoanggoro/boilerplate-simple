package app

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/app/factory"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/app/routes"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/config"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber  *fiber.App
	Config *config.Config
}

func NewApp() *App {
	container := factory.Build()
	app := fiber.New()
	a := &App{Fiber: app, Config: container.Config}

	// Register routes
	routes.NewRoutes(app, container)

	return a
}

func (a *App) Start() {
	configApp := a.Config.App
	logger.Infof("%s Starting server on : %s", configApp.Name, configApp.Port)
	a.Fiber.Listen(":" + configApp.Port)
}
