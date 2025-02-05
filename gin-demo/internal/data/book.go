package data

import "github.com/mukappalambda/go-examples/gin-demo/internal/core/repository"

var Books = []repository.Book{
	{
		ID:     1,
		Author: "alex",
		Title:  "alex's book",
	},
	{
		ID:     2,
		Author: "bob",
		Title:  "bob's book",
	},
}
