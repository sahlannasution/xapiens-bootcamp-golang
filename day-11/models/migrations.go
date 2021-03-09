package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {

	if check := db.Migrator().HasTable(&Post{}); !check {
		db.Migrator().CreateTable(&Post{})
		fmt.Println("Post Table has been created!")
	}

	if check := db.Migrator().HasTable(&Comment{}); !check {
		db.Migrator().CreateTable(&Comment{})
		fmt.Println("Comment Table has been created!")
	}

	if check := db.Migrator().HasTable(&Todos{}); !check {
		db.Migrator().CreateTable(&Todos{})
		fmt.Println("Todos Table has been created!")
	}

}
