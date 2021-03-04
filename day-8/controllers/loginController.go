package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type getBodyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Login func
func Login(c *gin.Context) {
	var body getBodyLogin
	err := c.Bind(&body)
	if err != nil {
		fmt.Println("Can't get data!")
	}

	// response
	c.JSON(200, gin.H{
		"status":       "success",
		"messages":     "Successfully Login.",
		"access_token": "sdnandssqiwnakjnklblksjnakdbiyBEIN",
		"role":         "user",
	})
}
