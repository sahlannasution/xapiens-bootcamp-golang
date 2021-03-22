package args

import "github.com/graphql-go/graphql"

func GetProductArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{ // untuk param
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
}

func CreateProductArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{ // untuk param
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"info": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
	}
}

func UpdateProductArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{ // untuk param
		"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"info": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
	}
}

func DeleteProductArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{ // untuk param
		"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
			Type: graphql.NewNonNull(graphql.Int),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"info": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"price": &graphql.ArgumentConfig{
			Type: graphql.Float,
		},
	}
}
