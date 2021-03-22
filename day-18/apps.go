package main

import (
	"net/http"
	"xapiens-bootcamp-golang/day-18/services"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	// routing "/" ibarat di rest seperti group
	route.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)
		result := services.ExecuteQuery(query.Query, services.Schema)
		c.JSON(http.StatusOK, result)
	})

	route.Run()
}
