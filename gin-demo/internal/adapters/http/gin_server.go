package http

import (
	"log"
	"net/http"

	httpPort "github.com/mukappalambda/go-examples/gin-demo/internal/core/ports/http"

	"github.com/gin-gonic/gin"
	"github.com/mukappalambda/go-examples/gin-demo/internal/core/repository"
	"github.com/mukappalambda/go-examples/gin-demo/internal/data"
)

type GinHttpServer struct {
	engine *gin.Engine
}

var _ httpPort.HTTPServer = (*GinHttpServer)(nil)

var books = data.Books

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func NewBook(c *gin.Context) {
	var book repository.Book

	if c.BindJSON(&book) == nil {
		log.Println(book.Id)
		log.Println(book.Author)
		log.Println(book.Title)
		books = append(books, book)
	}

	c.JSON(http.StatusCreated, book)
}

func NewServer() *GinHttpServer {
	engine := gin.Default()
	bookRoutes := engine.Group("/books")
	{
		bookRoutes.GET("/", GetBooks)
		bookRoutes.POST("/", NewBook)
	}
	srv := &GinHttpServer{engine: engine}
	return srv
}

func (s *GinHttpServer) Run(addr string) error {
	return http.ListenAndServe(addr, s.engine.Handler())
}
