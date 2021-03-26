package migrator

import (
	"fmt"
	"xapiens-bootcamp-golang/day-19/models"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {

	// Drop Table Users
	db.Migrator().DropTable(&models.Users{})
	// Drop Table Genres
	db.Migrator().DropTable(&models.Genres{})
	// Drop Table Movies
	db.Migrator().DropTable(&models.Movies{})
	// Drop Table MoviesGenres
	db.Migrator().DropTable(&models.MovieGenres{})

	// Drop Table MoviesGenres
	db.Migrator().DropTable(&models.Reviews{})

	if check := db.Migrator().HasTable(&models.Users{}); !check {
		db.Migrator().CreateTable(&models.Users{})
		fmt.Println("Table Users has been created!")
	}

	if check := db.Migrator().HasTable(&models.Genres{}); !check {
		db.Migrator().CreateTable(&models.Genres{})
		fmt.Println("Table Genres has been created!")
	}

	if check := db.Migrator().HasTable(&models.Movies{}); !check {
		db.Migrator().CreateTable(&models.Movies{})
		fmt.Println("Table Movies has been created!")
	}

	if check := db.Migrator().HasTable(&models.MovieGenres{}); !check {
		db.Migrator().CreateTable(&models.MovieGenres{})
		fmt.Println("Table MoviesGenres has been created!")
	}

	if check := db.Migrator().HasTable(&models.Reviews{}); !check {
		db.Migrator().CreateTable(&models.Reviews{})
		fmt.Println("Table Reviews has been created!")
	}
}
