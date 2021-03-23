package main

import (
	"net/http"
	"xapiens-bootcamp-golang/day-19/config"
	"xapiens-bootcamp-golang/day-19/migrator"
	"xapiens-bootcamp-golang/day-19/routes"
	"xapiens-bootcamp-golang/day-19/schema"
	"xapiens-bootcamp-golang/day-19/seeder"

	"github.com/gin-gonic/gin"
)

func main() {
	dbPG := config.Connection()     // db Connection
	migrator.Migrations(dbPG)       // migrate tables
	seeder.SeederUser(dbPG)         // seed Users Data
	seeder.SeederGenres(dbPG)       // seed Genres Data
	seeder.SeederMovies(dbPG)       // seed Movies Data
	seeder.SeederMoviesGenres(dbPG) // seed MoviesGenres Data
	seeder.SeederReview(dbPG)       // seed Reviews Data

	route := gin.Default()

	// Define route
	route.POST("/", func(c *gin.Context) {
		// Struvt Query
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)                                            // Get query params
		result := routes.ExecuteQuery(query.Query, schema.Schema) // Run Query
		c.JSON(http.StatusOK, result)                             // Send Response
	})

	route.Run()
}
