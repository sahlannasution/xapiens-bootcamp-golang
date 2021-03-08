package controllers

import (
	"net/http"
	"strconv"
	"xapiens-bootcamp-golang/day-10/models"

	"github.com/gin-gonic/gin"
)

// AddMoviesGenre func
func (StrDB *StrDB) AddMoviesGenre(c *gin.Context) {
	var (
		movieGenre models.MovieGenres
		result     gin.H
	)

	movieID, _ := strconv.ParseInt(c.PostForm("moviesID"), 10, 64)
	// genreID, _ := strconv.ParseInt(c.PostForm("genresID"), 10, 64)

	movieGenre.MoviesID = uint(movieID)
	// movieGenre.GenresID = uint(genreID)
	StrDB.DB.Create(&movieGenre)
	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Added Data!",
		"data":    movieGenre,
		// "data": map[string]interface{}{
		// 	"id":      movies.ID,
		// 	"title":   movies.Title,
		// 	"year":    movies.Year,
		// 	"ratings": movies.Ratings,
		// },
	}

	c.JSON(http.StatusOK, result)
}

//GetMoviesGenre func
// func (StrDB *StrDB) GetMoviesGenre(c *gin.Context) {
// 	var (
// 		movieGenre []models.MovieGenre
// 		result     gin.H
// 	)
// 	// MovieGenre struct
// 	type MovieGenre struct {
// 		MoviesID uint     `json:"movies_id"`
// 		GenresID uint     `json:"genres_id"`
// 		Genres   []Genres `gorm:"foreignKey:MovieGenreID" json:"Genres"`
// 		ID       uint     `gorm:"primarykey" json:"id"`
// 	}
// 	// Genres struct
// 	type Genres struct {
// 		Name string `json:"name"`
// 		ID   uint   `gorm:"primarykey" json:"id"`
// 	}

// 	// StrDB.DB.Find(&movies)
// 	// StrDB.DB.Joins("JOIN districts on districts.province_id=provinces.id")
// 	// // Joins("JOIN sub_districts on sub_districts.district_id=districts.id").
// 	result = gin.H{
// 		"status":  "success",
// 		"message": "Successfully Get Movies List",
// 		"data":    movieGenre,
// 	}

// 	c.JSON(http.StatusOK, result)
// }
