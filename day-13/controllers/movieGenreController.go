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
	genresID, _ := strconv.ParseInt(c.PostForm("genresID"), 10, 64)

	movieGenre.MoviesID = uint(movieID)
	movieGenre.GenresID = uint(genresID)
	StrDB.DB.Create(&movieGenre)
	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Added Data!",
		"data":    movieGenre,
	}

	c.JSON(http.StatusOK, result)
}
