package models

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeederMovie(db *gorm.DB) {
	var movieArray = [...][3]string{
		{"Parasite", "2019", "0"},
		{"The Avengers", "2017", "0"},
		{"Spiderman", "2010", "0"},
		{"Doraemon Stand By Me 2", "2020", "0"},
		{"Parasyte", "2016", "0"},
		{"Harry Potter", "2014", "0"},
		{"Ayat-Ayat Cinta", "2006", "0"},
	}

	var movie Movies

	for _, data := range movieArray {
		// Get Data from Array
		year, _ := strconv.ParseInt(data[1], 10, 64)
		rating, _ := strconv.ParseInt(data[2], 10, 64)
		movie.Title = data[0]
		movie.Year = int(year)
		movie.Ratings = int(rating)
		movie.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&movie)
	}
	fmt.Println("Movie Data has been Seed!")
}

func SeederGenre(db *gorm.DB) {
	var genreArray = [...]string{
		"Action",
		"Sci-Fi",
		"Mystery",
		"Comedy",
		"Drama",
		"Thriller",
		"Adventure",
		"Animation",
	}

	var genre Genres

	for _, data := range genreArray {
		// Get Data from Array
		genre.Name = data
		genre.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&genre)
	}
	fmt.Println("Genre Data has been Seed!")
}

func SeederUser(db *gorm.DB) {
	var userArray = [...][4]string{
		{"Sahlan.Nasution@xapiens.id", "admin@1234", "Sahlan Admin Xapiens", "admin"},
		{"sahlan.nasution07@gmail.com", "guest@1234", "Sahlan Nasution", "guest"},
		{"zahland.nasution@gmail.com", "guest@1234", "Sahlan NST", "guest"},
		{"armanzulfikri@mail.com", "guest@1234", "Arman Zulfikri", "guest"},
		{"wawan_setiawan@mail.com", "guest@1234", "Wawan Setiawan", "guest"},
		{"mas.dimas@mail.com", "guest@1234", "Mas Dimas", "guest"},
		{"dimas.maskun@mail.com", "guest@1234", "Dimas Chan wkwk", "guest"},
	}

	var user Users

	for _, data := range userArray {
		// Get Data from Array
		user.Email = data[0]
		user.Password = data[1]
		user.FullName = data[2]
		user.Role = data[3]

		// Encrypt Password using Bcrypt
		encrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		// Check if error while encrypting
		if err != nil {
			log.Println(err)
		}

		user.Password = string(encrypt)
		user.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&user)
	}
	fmt.Println("Users Data has been Seed!")
}
