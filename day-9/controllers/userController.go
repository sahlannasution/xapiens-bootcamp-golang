package controllers

import (
	"fmt"
	"net/http"
	"xapiens-bootcamp-golang/day-9/models"

	"github.com/gin-gonic/gin"
)

// Register func
func (StrDB *StrDB) Register(c *gin.Context) {
	var (
		users  models.Users
		result gin.H
	)
	// fmt.Println(c.Bind(&users))
	if err := c.Bind(&users); err != nil {
		fmt.Println("Error can't Regist the User!")
	} else {
		StrDB.DB.Create(&users)
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Register!",
			"data": map[string]interface{}{
				"id":       users.ID,
				"email":    users.Email,
				"fullName": users.FullName,
				"role":     users.Role,
			},
		}
	}

	c.JSON(http.StatusOK, result)
}

// Signin func
func (StrDB *StrDB) Signin(c *gin.Context) {
	var (
		users  models.Users
		result gin.H
	)

	// fmt.Println(c.Bind(&users))
	if err := c.Bind(&users); err != nil {
		fmt.Println("Error can't Login!")
	} else {
		email := c.PostForm("email")
		password := c.PostForm("password")
		StrDB.DB.Find(&users, "email = ? AND password = ?", email, password)
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Login!",
			"data":    users,
		}
	}

	c.JSON(http.StatusOK, result)
}

// func (StrDB *StrDB) GetVendorList(c *gin.Context) {
// 	var (
// 		vendor []models.Vendor
// 		result gin.H
// 	)

// 	StrDB.DB.Find(&vendor)
// 	result = gin.H{
// 		"msg":  "success",
// 		"data": vendor,
// 	}

// 	c.JSON(http.StatusOK, result)
// }

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
