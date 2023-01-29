package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Ramayan", Author: "Tulsidas", Quantity: 2},
	{ID: "2", Title: "Manikarnika", Author: "Vijendra Prasad", Quantity: 9},
	{ID: "3", Title: "Mahabharata", Author: "Maharaj Vyas", Quantity: 4},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newbook Book
	err := c.BindJSON(&newbook)
	if err != nil {
		return
	}
	books = append(books, newbook)
	c.IndentedJSON(http.StatusCreated, newbook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
