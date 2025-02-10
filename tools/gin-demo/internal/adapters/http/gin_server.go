package http

import (
	"log"
	"net/http"
	"time"

	httpPort "github.com/mukappalambda/go-examples/tools/gin-demo/internal/core/ports/http"

	"github.com/gin-gonic/gin"
	"github.com/mukappalambda/go-examples/tools/gin-demo/internal/core/repository"
	"github.com/mukappalambda/go-examples/tools/gin-demo/internal/data"
)

type GinHTTPServer struct {
	engine *gin.Engine
}

var _ httpPort.HTTPServer = (*GinHTTPServer)(nil)

var books = data.Books

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func NewBook(c *gin.Context) {
	var book repository.Book

	if c.BindJSON(&book) == nil {
		log.Println(book.ID)
		log.Println(book.Author)
		log.Println(book.Title)
		books = append(books, book)
	}

	c.JSON(http.StatusCreated, book)
}

func NewServer() *GinHTTPServer {
	engine := gin.Default()
	bookRoutes := engine.Group("/books")
	{
		bookRoutes.GET("/", GetBooks)
		bookRoutes.POST("/", NewBook)
	}
	srv := &GinHTTPServer{engine: engine}
	return srv
}

func (s *GinHTTPServer) Run(addr string) error {
	server := &http.Server{
		Addr:              addr,
		Handler:           s.engine.Handler(),
		ReadHeaderTimeout: 300 * time.Millisecond,
	}
	return server.ListenAndServe()
}
