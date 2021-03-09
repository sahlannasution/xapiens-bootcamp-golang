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

	// Post struct
	type Post struct {
		UserID uint   `json:"user_id"`
		ID     uint   `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	fmt.Println(body)
	fmt.Println("==============================")
	fmt.Println(string(body)) // convert body yang valuenya byte menjadi string
	// sb := string(body)        // convert body yang valuenya byte menjadi string
	data := Post{}

	// unmarshal dari hasil get api
	// unmarshall itu adalah convert dari tipe data .... ke bentuk JSON / object
	if error := json.Unmarshal(body, &data); error != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println(data)

	post.UserID = data.UserID
	post.ID = data.ID
	post.Title = data.Title
	post.Body = data.Body
	StrDB.DB.Create(&post)

	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Get and Insert Data!",
		"data":    post,
	}

	c.JSON(http.StatusOK, result)
}
