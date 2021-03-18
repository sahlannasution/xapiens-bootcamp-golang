package models

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeederReview(db *gorm.DB) {
	var reviewArray = [...][4]string{
		{"1", "4", "Keren Banget Filmnya", "9"},
		{"1", "5", "Keren Banget Filmnya", "8"},
		{"1", "6", "Bagus jalan cerita nya", "7"},
		{"2", "1", "sad ending", "3"},
		{"2", "2", "bagus tapi kurang nendang", "5"},
		{"2", "5", "siapa nih yang bikin karakter nya bagus banget", "8"},
		{"3", "1", "karakter utama nya gabagsu", "5"},
		{"3", "4", "ok", "6"},
		{"3", "6", "gooood", "7"},
	}

	var review Reviews

	for _, data := range reviewArray {
		// Get Data from Array
		user_id, _ := strconv.ParseInt(data[0], 10, 64)
		movie_id, _ := strconv.ParseInt(data[1], 10, 64)
		rate, _ := strconv.ParseInt(data[3], 10, 64)
		review.UsersID = uint(user_id)
		review.MoviesID = uint(movie_id)
		review.Review = data[2]
		review.Rate = int(rate)
		review.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&review)
	}
	fmt.Println("Review Data has been Seed!")
}

func SeederMovieGenre(db *gorm.DB) {
	var movieGenreArray = [...][2]string{
		{"1", "4"},
		{"1", "5"},
		{"1", "6"},
		{"2", "1"},
		{"2", "2"},
		{"2", "7"},
		{"3", "1"},
		{"3", "2"},
		{"3", "7"},
	}

	var movieGenre MovieGenres

	for _, data := range movieGenreArray {
		// Get Data from Array
		movie_id, _ := strconv.ParseInt(data[0], 10, 64)
		genre_id, _ := strconv.ParseInt(data[1], 10, 64)
		movieGenre.MoviesID = uint(movie_id)
		movieGenre.GenresID = uint(genre_id)
		movieGenre.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&movieGenre)
	}
	fmt.Println("Movie Genre Data has been Seed!")
}

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
	var userArray = [...][5]string{
		{"Sahlan.Nasution@xapiens.id", "admin@1234", "Sahlan Admin Xapiens", "admin", "1"},
		// {"sahlan.nasution07@gmail.com", "guest@1234", "Sahlan Nasution", "guest", "1"},
		// {"zahland.nasution@gmail.com", "guest@1234", "Sahlan NST", "guest", "1"},
		{"armanzulfikri@mail.com", "guest@1234", "Arman Zulfikri", "guest", "1"},
		{"wawan_setiawan@mail.com", "guest@1234", "Wawan Setiawan", "guest", "1"},
		{"mas.dimas@mail.com", "guest@1234", "Mas Dimas", "guest", "1"},
		{"dimas.maskun@mail.com", "guest@1234", "Dimas Chan wkwk", "guest", "1"},
	}

	var user Users

	for _, data := range userArray {
		// Get Data from Array
		user.Email = data[0]
		user.Password = data[1]
		user.FullName = data[2]
		user.Role = data[3]
		status, _ := strconv.ParseInt(data[4], 10, 64)
		user.Status = int(status)

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
