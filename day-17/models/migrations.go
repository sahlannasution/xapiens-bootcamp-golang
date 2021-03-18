package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {

	/* Drop Tables */

	/* Drop Tables */

	/* Check if table is in db postgre */

	db.Migrator().DropTable(&Users{})
	if check := db.Migrator().HasTable(&Users{}); !check {

		db.Migrator().CreateTable(&Users{})
		fmt.Println("Users Table has been created!")
	}
	db.Migrator().DropTable(&Genres{})
	if check := db.Migrator().HasTable(&Genres{}); !check {

		db.Migrator().CreateTable(&Genres{})
		fmt.Println("Genre Table has been created!")
	}

	if check := db.Migrator().HasTable(&Movies{}); !check {
		db.Migrator().DropTable(&Movies{})
		db.Migrator().CreateTable(&Movies{})
		fmt.Println("Movies Table has been created!")
	}
	db.Migrator().DropTable(&MovieGenres{})
	if check := db.Migrator().HasTable(&MovieGenres{}); !check {

		db.Migrator().CreateTable(&MovieGenres{})
		fmt.Println("Movies Genre Table has been created!")
	}
	db.Migrator().DropTable(&Reviews{})
	if check := db.Migrator().HasTable(&Reviews{}); !check {

		db.Migrator().CreateTable(&Reviews{})
		fmt.Println("Reviews Table has been created!")
	}

	db.Migrator().DropTable(&RegistMailer{})
	if check := db.Migrator().HasTable(&RegistMailer{}); !check {

		db.Migrator().CreateTable(&RegistMailer{})
		fmt.Println("Register Mailer Table has been created!")
	}
	/*  End Check if table is in db postgre */
}
