package controllers

import (
	"net/http"
	"strconv"
	logger "xapiens-bootcamp-golang/gogorm/log"
	"xapiens-bootcamp-golang/gogorm/models"

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

	movieGenre.MovieID = uint(movieID)
	movieGenre.GenreID = uint(genresID)
	if res := StrDB.DB.Create(&movieGenre); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Bad Request",
			"message": "Cant Process the Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		logger.Sentry(err)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Added Data!",
			"data":    movieGenre,
		}
		c.JSON(http.StatusOK, result)
	}
}
