package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	if check := db.Migrator().HasTable(&Users{}); !check {
		db.Migrator().CreateTable(&Users{})
		fmt.Println("Users Table has been created!")
	}
}
