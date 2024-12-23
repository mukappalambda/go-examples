package data

import "github.com/mukappalambda/go-examples/gin-demo/internal/core/repository"

var Books = []repository.Book{
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
