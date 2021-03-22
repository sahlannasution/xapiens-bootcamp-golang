package main

import (
	"net/http"
	"xapiens-bootcamp-golang/day-18/services"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Define route
	route.POST("/", func(c *gin.Context) {
		// Struvt Query
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)                                                // Get query params
		result := services.ExecuteQuery(query.Query, services.Schema) // Run Query
		c.JSON(http.StatusOK, result)                                 // Send Response
	})

	route.Run()
}
