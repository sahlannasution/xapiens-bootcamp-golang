package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	logger "xapiens-bootcamp-golang/day-15/log"
	"xapiens-bootcamp-golang/day-15/models"

	"github.com/gin-gonic/gin"
)

// AddReview func
func (StrDB *StrDB) AddReview(c *gin.Context) {
	var (
		reviews models.Reviews
		result  gin.H
	)
	userID, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	movieID, _ := strconv.ParseInt(c.PostForm("movie_id"), 10, 64)
	review := c.PostForm("review")
	rate, _ := strconv.ParseInt(c.PostForm("rate"), 10, 64)

	reviews.UsersID = uint(userID)
	reviews.MoviesID = uint(movieID)
	reviews.Review = review
	reviews.Rate = int(rate)

	if res := StrDB.DB.Create(&reviews); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Bad Request",
			"message": "Cant Process the Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		logger.Sentry(err)
	} else {
		fmt.Println(reviews)
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Add Reviews For this Movie!",
			"data":    reviews,
		}

		c.JSON(http.StatusOK, result)
	}

}

// GetReviewByMovie func
func (StrDB *StrDB) GetReviewByMovie(c *gin.Context) {

	var (
		reviews []models.Reviews
		result  gin.H
	)

	movieID := c.Param("movie_id")
	if res := StrDB.DB.Where("movies_id = ?", movieID).Find(&reviews); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Not Fpund!",
			"message": "Cant Find Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		logger.Sentry(err)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Successfully Get Review List",
			"data":    reviews,
		}

		c.JSON(http.StatusOK, result)
	}

}

// GetReviewDetail func
func (StrDB *StrDB) GetReviewDetail(c *gin.Context) {

	var (
		reviews models.Reviews
		result  gin.H
	)

	reviewID := c.Param("review_id")
	if res := StrDB.DB.Where("id = ?", reviewID).First(&reviews); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Not Fpund!",
			"message": "Cant Find Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		logger.Sentry(err)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Successfully Get Review Detail",
			"data":    reviews,
		}

		c.JSON(http.StatusOK, result)
	}
}
