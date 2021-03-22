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
			// URL : localhost:8080/?Query=mutation+_{createProduct(name:"sprite",info:"minuman soda",price:5000){id,name,info,price}}
			"createProduct": &graphql.Field{
				Type:    types.ProductType(),      // Define ProductType
				Args:    args.CreateProductArgs(), // Define Create Product Arguments (params to add)
				Resolve: resolver.CreateProduct,   // Call CreateProduct resolver function
			},

			// update data by id
			"updateProduct": &graphql.Field{
				Type:    types.ProductType(),      // Define ProductType
				Args:    args.UpdateProductArgs(), // Define Update Product Arguments (params to update)
				Resolve: resolver.UpdateProduct,   // Call UpdateProduct resolver function
			},

			"deleteProduct": &graphql.Field{
				Type:    types.ProductType(),      // Define ProductType
				Args:    args.DeleteProductArgs(), // Define Delete Product Arguments (params to add)
				Resolve: resolver.DeleteProduct,   // Call DeleteProduct resolver function
			},
		},
	},
)
