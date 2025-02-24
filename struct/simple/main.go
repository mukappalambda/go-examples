package main

import (
	"github.com/mukappalambda/go-examples/struct/simple/book"
	"github.com/mukappalambda/go-examples/struct/simple/library"
)

func main() {
	books := []*book.Book{
		book.New(1, "alpha", "alpha's book"),
		book.New(2, "beta", "beta's book"),
		book.New(3, "gamma", "gamma's book"),
	}
	library := library.New(books)
	library.DisplayBooks()

	firstBook := book.New(4, "", "delta's book")
	secondBook := book.New(5, "epsilon", "")
	firstBook.SetAuthor("delta")
	secondBook.SetTitle("epsilon's book")

	otherBooks := []*book.Book{firstBook, secondBook}
	library.AddBooks(otherBooks...)
	library.DisplayBooks()
}
