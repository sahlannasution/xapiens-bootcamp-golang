package services

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// Define schema for route
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,    // Define query type for get data
		Mutation: mutationType, // Define mutation type for create, update, delete
	},
)

// func ExecuteQuery
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// if error
	if len(result.Errors) > 0 {
		fmt.Println("ada error : ", result.Errors)
	}

	return result
}
