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

	// Todos struct
	type Todos struct {
		UserID    uint   `json:"user_id"`
		ID        uint   `gorm:"primarykey" json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	fmt.Println(body)
	fmt.Println("==============================")
	fmt.Println(string(body)) // convert body yang valuenya byte menjadi string
	// sb := string(body)        // convert body yang valuenya byte menjadi string
	data := Todos{}

	// unmarshal dari hasil get api
	// unmarshall itu adalah convert dari tipe data .... ke bentuk JSON / object
	if error := json.Unmarshal(body, &data); error != nil {
		fmt.Println("Error ", err.Error())
	}
	fmt.Println(data)

	todos.UserID = data.UserID
	todos.ID = data.ID
	todos.Title = data.Title
	todos.Completed = data.Completed
	StrDB.DB.Create(&todos)

	result = gin.H{
		"status":  "success",
		"message": "Sucessfully Get and Insert Data!",
		"data":    todos,
	}

	c.JSON(http.StatusOK, result)
}
