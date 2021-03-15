package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {

	/* Drop Tables */
	db.Migrator().DropTable(&Users{})
	db.Migrator().DropTable(&Genres{})
	db.Migrator().DropTable(&Movies{})
	/* Drop Tables */

	/* Check if table is in db postgre */
	if check := db.Migrator().HasTable(&Users{}); !check {
		db.Migrator().CreateTable(&Users{})
		fmt.Println("Users Table has been created!")
	}

	if check := db.Migrator().HasTable(&Movies{}); !check {
		db.Migrator().CreateTable(&Movies{})
		fmt.Println("Movies Table has been created!")
	}

	if check := db.Migrator().HasTable(&Genres{}); !check {
		db.Migrator().CreateTable(&Genres{})
		fmt.Println("Genre Table has been created!")
	}

	if check := db.Migrator().HasTable(&MovieGenres{}); !check {
		db.Migrator().CreateTable(&MovieGenres{})
		fmt.Println("Movies Genre Table has been created!")
	}

	if check := db.Migrator().HasTable(&Reviews{}); !check {
		db.Migrator().CreateTable(&Reviews{})
		fmt.Println("Reviews Table has been created!")
	}
	/*  End Check if table is in db postgre */
}
