package factory

import (
	"os"

	"github.com/alfianvitoanggoro/boilerplate-simple/internal/config"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/delivery"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/db"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/migration"
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/infrastructure/repository"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/logger"
	"gorm.io/gorm"
)

type Container struct {
	Config      *config.Config
	UserHandler *delivery.UserHandler
	AuthHandler *delivery.AuthHandler
	DB          *gorm.DB
}

func Build() *Container {
	cfg := config.Load()

	database, err := db.ConnectDB(cfg.DB)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err.Error())
	}

	// Migration database if MIGRATE env is true
	args := os.Args
	if len(args) > 1 && args[1] == "migrate" {
		migration.Run(database)
		os.Exit(0)
	}

	// Register Factories
	userRepo := repository.NewUserRepository(database)
	userHandler := newUserFactory(userRepo)
	authHandler := newAuthFactory(userRepo)

	return &Container{
		Config:      cfg,
		UserHandler: userHandler,
		AuthHandler: authHandler,
		DB:          database,
	}
}
