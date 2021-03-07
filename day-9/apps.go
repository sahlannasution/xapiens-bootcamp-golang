package main

import (

	// "xapiens-day9/controller"
	// "xapiens-day9/models"

	"xapiens-bootcamp-golang/day-9/config"
	"xapiens-bootcamp-golang/day-9/controllers"
	"xapiens-bootcamp-golang/day-9/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Db Connections
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}

	models.Migrations(dbPG)

	routing := gin.Default()

	//routes

	routing.POST("/register", strDB.Register)
	routing.POST("/signin", strDB.Signin)
	// routing.GET("/vendorList", strDB.GetVendorList)
	// routing.GET("/VendorByQuery", strDB.GetDataVendor)
	routing.Run()
}
