package services

import (
	"xapiens-bootcamp-golang/day-18/args"
	"xapiens-bootcamp-golang/day-18/resolver"
	"xapiens-bootcamp-golang/day-18/types"

	"github.com/graphql-go/graphql"
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			// urlnya untuk create data
			// localhost:8080/?Query=mutation+_{create(name:"sprite",info:"minuman soda",price:5000){id,name,info,price}}
			"createProduct": &graphql.Field{
				Type: types.ProductType(),
				Args: args.CreateProductArgs(),
				// kalau di rest resolve itu kayak controller
				Resolve: resolver.CreateProduct,
			},

			// update data by id
			"updateProduct": &graphql.Field{
				Type: types.ProductType(),
				Args: args.UpdateProductArgs(),
				// kalau di rest resolve itu kayak controller
				Resolve: resolver.UpdateProduct,
			},

			"deleteProduct": &graphql.Field{
				Type: types.ProductType(),
				Args: args.DeleteProductArgs(),
				// kalau di rest resolve itu kayak controller
				Resolve: resolver.DeleteProduct,
			},
		},
	},
)
