package resolver

import (
	"log"
	"math/rand"
	"time"
	"xapiens-bootcamp-golang/day-18/models"

	"github.com/graphql-go/graphql"
)

// var product = models.
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

func GetProductList(params graphql.ResolveParams) (interface{}, error) {
	return Product, nil
}

func GetProduct(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		for _, products := range Product { // looping data dari slice product
			// lakukan pengecekan jika sama id di args == id di slice product
			// maka tampilkan data
			if int(products.ID) == id {
				return products, nil
			}
		}
	}
	return nil, nil
}

func CreateProduct(p graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano()) // ini untuk generate random value

	// bikin variable untuk parse data ke struct Product
	add := models.ProductModels{
		ID:    int64(rand.Intn(1000)),
		Name:  p.Args["name"].(string),
		Info:  p.Args["info"].(string),
		Price: p.Args["price"].(float64),
	}
	Product = append(Product, add) //ini untuk nambah data di database (var product)
	return add, nil                // untuk response yang akan ditampilkan
}

func UpdateProduct(p graphql.ResolveParams) (interface{}, error) {
	// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
	id := p.Args["id"].(int)
	name, checkName := p.Args["name"].(string)
	info, checkInfo := p.Args["info"].(string)
	price, checkPrice := p.Args["price"].(float64)

	log.Println("ini argsnyaa......", p.Args["name"].(string))

	productVar := models.ProductModels{}

	for i, v := range Product { // product adalah database
		if id == int(v.ID) { // ketika id di database sama dengan args
			if checkName {
				Product[i].Name = name
			}
			if checkInfo {
				Product[i].Info = info
			}
			if checkPrice {
				Product[i].Price = price
			}
			productVar = Product[i]
			break
		}
	}
	return productVar, nil // untuk response yang akan ditampilkan
}

func DeleteProduct(p graphql.ResolveParams) (interface{}, error) {
	// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
	id := p.Args["id"].(int)

	productVar := models.ProductModels{}

	for i, v := range Product { // product adalah database
		if id == int(v.ID) { // kalau data id dari database == data id dari args
			productVar = Product[i]
			// untuk hapus bisa menggunakan slice
			Product = append(Product[:i], Product[i+1:]...) // ini untuk hapus data dengan penerapan slice
			log.Println("ini Productnya....", Product)
		}
	}
	return productVar, nil // untuk response yang akan ditampilkan
}
