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
			// get all data product
			// url nya localhost:8080/?Query={list{id, name, info, price}}
			"getProductList": &graphql.Field{
				Type: graphql.NewList(types.ProductType()),
				// kalau di rest resolve itu kayak controller
				Resolve: resolver.GetProductList,
			},

			// get data with param
			// url untuk akses endpointnya localhost:8080/?Query={product(id:1){id, info, name, price}}
			"getProduct": &graphql.Field{
				Type: types.ProductType(),
				Args: args.GetProductArgs(),
				// kalau di rest resolve itu kayak controller
				Resolve: resolver.GetProduct,
			},
		},
	},
)
