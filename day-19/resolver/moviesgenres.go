package resolver

import (
	"xapiens-bootcamp-golang/day-19/config"
	"xapiens-bootcamp-golang/day-19/models"

	"github.com/graphql-go/graphql"
)

// func AddMovies
func AddMoviesGenres(params graphql.ResolveParams) (interface{}, error) {
	movies_id, checkMovies := params.Args["movies_id"].(int) // Get movies_id from params arguments
	genres_id, checkGenres := params.Args["genres_id"].(int) // Get genres_id from params arguments
	dbPG := config.Connection()
	// movies struct
	var (
		moviesGenres models.MovieGenres
	)
	if checkMovies && checkGenres {
		moviesGenres.MoviesID = uint(movies_id)
		moviesGenres.GenresID = uint(genres_id)
		dbPG.Create(&moviesGenres)
		return moviesGenres, nil //return movies data response
	}

	return nil, nil //return movies data response
}
