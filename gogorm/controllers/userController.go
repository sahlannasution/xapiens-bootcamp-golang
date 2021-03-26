package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	logger "xapiens-bootcamp-golang/gogorm/log"
	"xapiens-bootcamp-golang/gogorm/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Users struct
type Users struct {
	ID       string `gorm:"primarykey" json:"id"`
	Email    string `gorm:"primarykey" json:"email"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	Status   int    `json:"status"`
}

// Register func
func (StrDB *StrDB) Register(c *gin.Context) {
	var (
		users  models.Users
		mailer models.RegistMailer
		result gin.H
	)
	if err := c.Bind(&users); err != nil || users.Email == "" || users.Password == "" || users.FullName == "" || users.Role == "" {
		e := "Field Email, Password, FullName, Role is required!"
		result = gin.H{
			"status":  "bad request",
			"message": e,
		}
		fmt.Println("Field Email, Password, FullName, Role is required!")
		c.JSON(http.StatusBadRequest, result)

		logger.Sentry(err) // push log error ke sentry
	} else {
		// mail := users.Email
		// fmt.Println(mail)
		encrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			logger.Sentry(err)
		}
		users.Password = string(encrypt)
		if res := StrDB.DB.Create(&users); res.Error != nil {
			err := res.Error
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "bad request",
				"message": err.Error(),
			})

			logger.Sentry(err) // push log error ke sentry
		} else {
			mail := users.Email
			mailer.Type = "Send Email"
			mailer.Email = mail
			mailer.Status = false
			mailer.Message = fmt.Sprintf("Hello %s, your account has been successfully registered!", users.FullName)

			StrDB.DB.Create(&mailer)
			// fmt.Println(d)
			// gocron.Every(10).Seconds().Do("")

			// <-gocron.Start()
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
			c.JSON(http.StatusOK, result)
		}
	}
}

// Signin func
func (StrDB *StrDB) Signin(c *gin.Context) {
	var (
		users  Users
		result gin.H
	)
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		err := "Email and Password Cannot be Empty!"
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": err,
		})
		logger.SentryStr(err)
	} else {
		if res := StrDB.DB.Where("email = ? AND  password = ?", email, password).First(&users); res.Error != nil {
			err := res.Error
			result = gin.H{
				"status":  "not found",
				"message": "User not found!",
				"errors":  err.Error(),
			}
			c.JSON(http.StatusNotFound, result)
			logger.Sentry(err)
		} else {
			fmt.Println(res.Error)
			result = gin.H{
				"status":  "success",
				"message": "Sucessfully Login!",
				"data":    users,
			}
			c.JSON(http.StatusOK, result)
		}
	}

}

// GetUserData func
func (StrDB *StrDB) GetUserData(c *gin.Context) {
	var (
		users  Users
		result gin.H
	)
	email := c.Param("email")

	if email == "" {
		err := "Email and Password Cannot be Empty!"
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": err,
		})
		logger.SentryStr(err)
	} else {
		if check, data := CheckRedis(email); check {
			if err := json.Unmarshal(data, &users); err != nil {
				fmt.Println("Error", err.Error())
			}
			fmt.Println(users)
			result = gin.H{
				"status":  "success",
				"message": "Sucessfully Get Data!",
				"data":    users,
			}
			c.JSON(http.StatusOK, result)

		} else {
			if res := StrDB.DB.Where("email = ?", email).First(&users); res.Error != nil {
				err := res.Error
				result = gin.H{
					"status":  "not found",
					"message": "User not found!",
					"errors":  err.Error(),
				}
				c.JSON(http.StatusNotFound, result)
				logger.Sentry(err)
			} else {
				// fmt.Println(res.Error)
				output, _ := json.Marshal(users)

				result = gin.H{
					"status":  "success",
					"message": "Sucessfully Get Data!",
					"data":    users,
				}
				c.JSON(http.StatusOK, result)
				SetRedis(email, string(output))
			}
		}

	}
}

// UpdateUser func
func (StrDB *StrDB) UpdateUser(c *gin.Context) {
	var (
		users  Users
		result gin.H
	)
	email := c.Param("email")
	fullName := c.PostForm("fullName")

	if email == "" || fullName == "" {
		err := "Email and FullName Cannot be Empty!"
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": err,
		})
		// logger.Sentry(err)
	} else {
		if res := StrDB.DB.Where("email = ?", email).First(&users); res.Error != nil {
			err := res.Error
			result = gin.H{
				"status":  "not found",
				"message": "User not found!",
				"errors":  err.Error(),
			}
			c.JSON(http.StatusNotFound, result)
			logger.Sentry(err)
		} else {
			users.FullName = fullName
			StrDB.DB.Save(&users)
			result = gin.H{
				"status":  "success",
				"message": "Sucessfully Updated Data!",
				"data":    users,
			}

			c.JSON(http.StatusOK, result)
		}
	}
	// StrDB.DB.Find(&users, "email = ? AND password = ?", email, password)
	// StrDB.DB.Where("email = ?", email).Find(&users)

}
