package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	logger "xapiens-bootcamp-golang/day-13/log"

	"github.com/gin-gonic/gin"
)

// AddReview func
func (StrDB *StrDB) AddReview(c *gin.Context) {

	// Users struct
	type Users struct {
		ID       string `gorm:"primarykey" json:"id"`
		Email    string `gorm:"primarykey" json:"email"`
		FullName string `json:"fullName"`
		Role     string `json:"role"`
	}
	// // Reviews struct
	type Reviews struct {
		ID      uint   `gorm:"primarykey, autoIncrement" json:"id"`
		Review  string `json:"review"`
		Rate    int    `json:"rate"`
		UserID  int    `json:"user_id"`
		Users   Users  `gorm:"foreignKey:ID" json:"user"`
		MovieID int    `json:"movie_id"`
		Movies  Movies `gorm:"foreignKey:ID" json:"movie"`
	}
	var (
		reviews Reviews
		result  gin.H
	)
	userID, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	movieID, _ := strconv.ParseInt(c.PostForm("movie_id"), 10, 64)
	review := c.PostForm("review")
	rate, _ := strconv.ParseInt(c.PostForm("rate"), 10, 64)

	reviews.UserID = int(userID)
	reviews.MovieID = int(movieID)
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
		StrDB.DB.Preload("Users").Preload("Movies").First(&reviews)
		// StrDB.DB.Table("users").Select("users.full_name, users.email").Joins("Users").Joins("Movies").Find(&reviews)
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Add Reviews For this Movie!",
			"data": map[string]interface{}{
				"id":     reviews.ID,
				"movie":  reviews.Movies,
				"user":   reviews.Users,
				"review": reviews.Review,
				"rate":   reviews.Rate,
			},
		}

		c.JSON(http.StatusOK, result)
	}

}

// GetReviewByMovie func
func (StrDB *StrDB) GetReviewByMovie(c *gin.Context) {

	type Users struct {
		ID       string `gorm:"primarykey" json:"id"`
		Email    string `gorm:"primarykey" json:"email"`
		FullName string `json:"fullName"`
		Role     string `json:"role"`
	}

	type Reviews struct {
		ID     uint   `gorm:"primarykey, autoIncrement" json:"id"`
		Review string `json:"review"`
		Rate   int    `json:"rate"`
		UserID int    `json:"user_id"`
		Users  Users  `gorm:"foreignKey:ID" json:"user"`
		// MovieID uint    `json:"movie_id"`
		// Movies  Movies  `gorm:"foreignKey:ID" json:"movie"`
	}

	var (
		reviews []Reviews
		result  gin.H
	)

	movieID := c.Param("movie_id")
	if res := StrDB.DB.Preload("Users").Where("movie_id = ?", movieID).Find(&reviews); res.Error != nil {
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
	// StrDB.DB.Preload("Users").Find(&reviews, "movie_id = ?", movieID)

}

// GetReviewDetail func
func (StrDB *StrDB) GetReviewDetail(c *gin.Context) {

	type Users struct {
		ID       string `gorm:"primarykey" json:"id"`
		Email    string `gorm:"primarykey" json:"email"`
		FullName string `json:"fullName"`
		Role     string `json:"role"`
	}

	type Reviews struct {
		ID     uint   `gorm:"primarykey, autoIncrement" json:"id"`
		Review string `json:"review"`
		Rate   int    `json:"rate"`
		// UserID uint   `json:"user_id"`
		Users Users `gorm:"foreignKey:ID" json:"user"`
		// MovieID uint    `json:"movie_id"`
		Movies Movies `gorm:"foreignKey:ID" json:"movie"`
	}

	var (
		reviews Reviews
		result  gin.H
	)

	reviewID := c.Param("review_id")
	if res := StrDB.DB.Preload("Movies").Where("id = ?", reviewID).Find(&reviews); res.Error != nil {
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
