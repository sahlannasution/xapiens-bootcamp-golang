package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

	StrDB.DB.Create(&reviews)
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
	StrDB.DB.Preload("Users").Find(&reviews, "movie_id = ?", movieID)
	result = gin.H{
		"status":  "success",
		"message": "Successfully Get Review List",
		"data":    reviews,
	}

	c.JSON(http.StatusOK, result)
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
	StrDB.DB.Preload("Users").Preload("Movies").Find(&reviews, "id = ?", reviewID)
	result = gin.H{
		"status":  "success",
		"message": "Successfully Get Review List",
		"data":    reviews,
	}

	c.JSON(http.StatusOK, result)
}
