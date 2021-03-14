package controllers

import (
	"net/http"
	"strconv"
	logger "xapiens-bootcamp-golang/day-13/log"
	"xapiens-bootcamp-golang/day-13/models"

	"github.com/gin-gonic/gin"
)

// Movies struct
type Movies struct {
	ID      string `gorm:"primarykey" json:"id"`
	Title   string `json:"title"`
	Year    string `json:"year"`
	Ratings string `json:"ratings"`
	// gorm.Model
}

// AddMovies func
func (StrDB *StrDB) AddMovies(c *gin.Context) {
	var (
		movies models.Movies
		result gin.H
	)
	title := c.PostForm("title")
	year, _ := strconv.ParseInt(c.PostForm("year"), 10, 64)
	ratings, _ := strconv.ParseInt(c.PostForm("ratings"), 10, 64)

	movies.Title = title
	movies.Year = int(year)
	movies.Ratings = int(ratings)
	if res := StrDB.DB.Create(&movies); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Internal Server Error",
			"message": "Cant create Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		logger.Sentry(err)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Created Data!",
			"data": map[string]interface{}{
				"id":      movies.ID,
				"title":   movies.Title,
				"year":    movies.Year,
				"ratings": movies.Ratings,
			},
		}
		c.JSON(http.StatusOK, result)
	}

}

// GetMoviesList func
func (StrDB *StrDB) GetMoviesList(c *gin.Context) {
	var (
		movies []models.Movies
		result gin.H
	)

	if res := StrDB.DB.Preload("Genres").Find(&movies); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Not Found",
			"message": "Cannot Get Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		logger.Sentry(err)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Get Data!",
			"data":    movies,
		}
		c.JSON(http.StatusOK, result)
	}
}
