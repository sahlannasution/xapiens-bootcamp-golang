package controllers

import (
	"net/http"
	"strconv"
	"xapiens-bootcamp-golang/day-10/models"

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
	StrDB.DB.Create(&movies)
	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Register!",
		"data": map[string]interface{}{
			"id":      movies.ID,
			"title":   movies.Title,
			"year":    movies.Year,
			"ratings": movies.Ratings,
		},
	}

	c.JSON(http.StatusOK, result)
}

// GetMoviesList func
func (StrDB *StrDB) GetMoviesList(c *gin.Context) {
	var (
		movies []Movies
		result gin.H
	)

	StrDB.DB.Find(&movies)
	result = gin.H{
		"status":  "success",
		"message": "Successfully Get Movies List",
		"data":    movies,
	}

	c.JSON(http.StatusOK, result)
}

// // GetDataVendor
// func (StrDB *StrDB) GetDataVendor(c *gin.Context) {
// 	var (
// 		vendor []models.Vendor
// 		result gin.H
// 	)

// 	StrDB.DB.Find(&vendor, "vendor_id = ?", c.Query("vendorID"))
// 	result = gin.H{
// 		"msg":  "success",
// 		"data": vendor,
// 	}

// 	c.JSON(http.StatusOK, result)
// }
