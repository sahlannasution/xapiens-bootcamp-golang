package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xapiens-bootcamp-golang/day-11/models"

	"github.com/gin-gonic/gin"
)

// Todos func
func (StrDB *StrDB) Todos(c *gin.Context) {
	var (
		todos  models.Todos
		result gin.H
	)
	// ambil data dari api
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

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
	if error := json.Unmarshal(body, &todos); error != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println(todos)

	StrDB.DB.Create(&todos)

	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Get and Insert Data!",
		"data":    todos,
	}

	c.JSON(http.StatusOK, result)
}
