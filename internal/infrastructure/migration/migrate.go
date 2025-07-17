package migration

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/internal/domain"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/logger"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		logger.Error("Migration failed: ", err)
		return err
	}
	logger.Info("Migration success!")
	return nil
}
