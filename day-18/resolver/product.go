package resolver

import (
	"math/rand"
	"time"
	"xapiens-bootcamp-golang/day-18/models"

	"github.com/graphql-go/graphql"
)

// Define slice Product Data from ProductModels
var Product = []models.ProductModels{
	{
		ID:    1,
		Name:  "Kopi luwak",
		Info:  "Nyaman di lambung nggak bikin kembung",
		Price: 1500,
	},
	{
		ID:    2,
		Name:  "Indomie",
		Info:  "Mie paling mantap",
		Price: 2500,
	},
	{
		ID:    3,
		Name:  "Coki coki",
		Info:  "Coklat asli",
		Price: 500,
	},
}

// func GetProductList
func GetProductList(params graphql.ResolveParams) (interface{}, error) {
	return Product, nil //return Product slice
}

func GetProduct(params graphql.ResolveParams) (interface{}, error) {

	id, ok := params.Args["id"].(int) // Get id from params arguments
	// if ok, find data in slice Product
	if ok {
		for _, products := range Product {
			if int(products.ID) == id { // if id find in slice
				return products, nil // return product detail
			}
		}
	}
	return nil, nil // else return no data
}

func CreateProduct(p graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano()) // generate random value for id

	// create variable using models ProductModels and define the value
	add := models.ProductModels{
		ID:    int64(rand.Intn(1000)),
		Name:  p.Args["name"].(string),
		Info:  p.Args["info"].(string),
		Price: p.Args["price"].(float64),
	}
	Product = append(Product, add) // append the struct to slice Product
	return add, nil                // send new data response
}

func UpdateProduct(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(int)                       // get id from params arguments
	name, checkName := p.Args["name"].(string)     // get name from params arguments
	info, checkInfo := p.Args["info"].(string)     // get info from params arguments
	price, checkPrice := p.Args["price"].(float64) // get price from params arguments

	productVar := models.ProductModels{} // Define productVar to patch newData

	for i, v := range Product { // find data
		if id == int(v.ID) { // if id finded
			if checkName {
				Product[i].Name = name // change data
			}
			if checkInfo {
				Product[i].Info = info // change data
			}
			if checkPrice {
				Product[i].Price = price // change data
			}
			productVar = Product[i] // patch data to productVar
			break                   // break loop
		}
	}
	return productVar, nil // send new data to response
}

func DeleteProduct(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(int) // get id from params arguments

	productVar := models.ProductModels{} // define struct to patch deleted data

	for i, v := range Product { // find in slice Product
		if id == int(v.ID) { // if id finded do this
			productVar = Product[i]
			// remove data from slice
			Product = append(Product[:i], Product[i+1:]...)
		}
	}
	return productVar, nil // send deleted data to response
}
