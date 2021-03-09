package main

import (
	// "xapiens-bootcamp-golang-day-11/config"
	// "xapiens-bootcamp-golang-day-11/controllers"
	// "xapiens-bootcamp-golang-day-11/models"

	"xapiens-bootcamp-golang/day-11/config"
	"xapiens-bootcamp-golang/day-11/controllers"
	"xapiens-bootcamp-golang/day-11/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// koneksi ke database postgree
	dbPG := config.Connection()
	strDB := controllers.StrDB{DB: dbPG}

	// code untuk migrasi
	models.Migrations(dbPG)

	// routing end point
	routing := gin.Default()

	// // routing vendor
	// routing.POST("/vendor", strDB.PostDataVendor)
	// routing.GET("/vendorList", strDB.GetVendorList)
	// routing.GET("/VendorByQuery", strDB.GetDataVendor)

	// // routing employee
	// routing.POST("/employee", strDB.PostDataEmployee)
	// routing.GET("/employeeList", strDB.GetEmployeeList)

	// // routing upload files
	// routing.POST("/uploadFile", controller.UploadSingleFile)
	// routing.POST("/uploadMultipleFile", controller.UploadMultipleFile)

	//routing 3rd api / http request
	routing.GET("/post", strDB.Post)
	routing.GET("/comment", strDB.Comment)
	routing.GET("/todos", strDB.Todos)
	routing.Run()
}
