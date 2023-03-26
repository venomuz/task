package migration

import (
	"github.com/venomuz/alif-task/internal/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {

	err := db.AutoMigrate(
		&models.Users{},
	)

	return err
}
