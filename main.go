package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/books", func(c *gin.Context) {
		log.Println("Handling GET /books request")
		c.JSON(http.StatusOK, books)
	})
	return r
}

// main is the entry point of the program.
//
// It initializes a new Gin router and defines a GET route for "/books".
// This route returns a JSON response with the HTTP status code 200 and the
// list of books.
//
// No parameters are required.
// No return values.
func main() {
	r := setupRouter()
	log.Println("Starting server on port 8080")
	r.Run(":8080")
}
