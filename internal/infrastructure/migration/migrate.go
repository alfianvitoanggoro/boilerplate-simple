package migration

import (
	user_model "github.com/alfianvitoanggoro/boilerplate-simple/internal/domain/user/model"
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/logger"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	if err := db.AutoMigrate(&user_model.User{}); err != nil {
		logger.Error("Migration failed: ", err)
		return err
	}
	logger.Info("Migration success!")
	return nil
}
