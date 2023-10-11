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
	// Health check route
	r.GET("/hc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	return r
}

// @Go Gin API
// @version 1.0
// @description This is a sample server for Your API.
// @host localhost:8080
// @BasePath /
func main() {
	r := setupRouter()
	log.Println("Starting server on port 8080")
	r.Run(":8080")
}
