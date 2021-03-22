package services

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// semacam routing di rest
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,    // query itu hanya untuk get data
		Mutation: mutationType, //ini untuk Create Update Delete
	},
)

// function untuk execute query
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// check error condition saat run var result
	if len(result.Errors) > 0 {
		fmt.Println("ada error : ", result.Errors)
	}

	return result
}
