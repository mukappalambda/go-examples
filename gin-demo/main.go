package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int32  `form:"id" json:"id,omitempty"`
	Author string `form:"author" json:"author,omitempty"`
	Title  string `form:"title" json:"title,omitempty"`
}

var books = []Book{
	{
		Id:     1,
		Author: "alex",
		Title:  "alex's book",
	},
	{
		Id:     2,
		Author: "bob",
		Title:  "bob's book",
	},
}

func main() {
	r := gin.Default()
	BookRoutes := r.Group("/books")
	{
		BookRoutes.GET("/", GetBooks)
		BookRoutes.POST("/", NewBook)
	}

	r.Run()
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func NewBook(c *gin.Context) {
	var book Book

	if c.BindJSON(&book) == nil {
		log.Println(book.Id)
		log.Println(book.Author)
		log.Println(book.Title)
		books = append(books, book)
	}

	c.JSON(http.StatusCreated, book)
}
