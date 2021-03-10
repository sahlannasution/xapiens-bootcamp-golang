package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xapiens-bootcamp-golang/day-11/models"

	"github.com/gin-gonic/gin"
)

// Post func
func (StrDB *StrDB) Post(c *gin.Context) {
	var (
		post   models.Post
		result gin.H
	)
	// ambil data dari api
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	// check kondisi saat ambil data dari api error / tidak
	if err != nil { // ini kalau error
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	// check kondisi saat membaca data body dari api
	if err != nil { // ini kalau error
		fmt.Println(err.Error())
	}

	fmt.Println(body)
	fmt.Println("==============================")
	fmt.Println(string(body)) // convert body yang valuenya byte menjadi string
	// sb := string(body)        // convert body yang valuenya byte menjadi string

	// unmarshal dari hasil get api
	// unmarshall itu adalah convert dari tipe data .... ke bentuk JSON / object
	if error := json.Unmarshal(body, &post); error != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println(post)
	StrDB.DB.Create(&post)

	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Get and Insert Data!",
		"data":    post,
	}

	c.JSON(http.StatusOK, result)
}
