package models

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		User{},
		Post{},
		Like{},
	}
	return db.AutoMigrate(models...)
}
