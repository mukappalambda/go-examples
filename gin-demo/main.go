package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Book struct {
  Id     int32    `form:"id"`
  Author string `form:"author"`
  Title  string `form:"title"`
}

var books = []Book{
  {
    Id: 1,
    Author: "alex",
    Title: "alex's book",
  },
  {
    Id: 2,
    Author: "bob",
    Title: "bob's book",
  },
}

func main()  {
  r := gin.Default()
  BookRoutes := r.Group("/books")
  {
    BookRoutes.GET("/", GetBooks)
    BookRoutes.POST("/", NewBook)
  }

  r.Run()
}

func GetBooks(c *gin.Context)  {
  c.JSON(200, books)
}

func NewBook(c *gin.Context) {
  var book Book

  if c.BindJSON(&book) == nil {
    log.Println(book.Id)
    log.Println(book.Author)
    log.Println(book.Title)
    books = append(books, book)
  }

  c.JSON(201, book)
}