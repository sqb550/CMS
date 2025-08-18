package database

import (
	"gorm.io/gorm"

	"CMS/app/models"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.ReportedPost{},
	)
	return err
}
