package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xapiens-bootcamp-golang/day-12/models"

	"github.com/gin-gonic/gin"
)

// Comment func
func (StrDB *StrDB) Comment(c *gin.Context) {
	var (
		comment models.Comment
		result  gin.H
	)
	// ambil data dari api
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments/1")

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
	if error := json.Unmarshal(body, &comment); error != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println(comment)

	StrDB.DB.Create(&comment)

	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Get and Insert Data!",
		"data":    comment,
	}

	c.JSON(http.StatusOK, result)
}
