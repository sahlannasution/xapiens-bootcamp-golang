package services

import (
	"xapiens-bootcamp-golang/day-18/args"
	"xapiens-bootcamp-golang/day-18/resolver"
	"xapiens-bootcamp-golang/day-18/types"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Get All Product Data
			// URL : localhost:8080/?Query={getProductList{id, name, info, price}}
			"getProductList": &graphql.Field{
				Type:    graphql.NewList(types.ProductType()), // Define Product Type
				Resolve: resolver.GetProductList,              // Call GetProducList resolver function
			},

			// Get Product By Id
			// url localhost:8080/?Query={getProduct(id:1){id, info, name, price}}
			"getProduct": &graphql.Field{
				Type:    types.ProductType(),   // Define Product Type
				Args:    args.GetProductArgs(), // Define Arguments to getProduct
				Resolve: resolver.GetProduct,   // Call GetProduct resolver function
			},
		},
	},
)
