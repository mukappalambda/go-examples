package main

import "fmt"

type Book struct {
	ID     int32
	Author string
	Title  string
}

var books = []Book{
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
	{
		ID:     3,
		Author: "joe",
		Title:  "joe's book",
	},
}

func main() {
	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}

	var oldBook Book
	for _, book := range books {
		if book.Author == "alex" {
			oldBook.ID = book.ID
			oldBook.Author = book.Author
			oldBook.Title = book.Title
		}
	}

	fmt.Printf("oldBook: %+v\n", oldBook)

	firstBook := books[0]
	fmt.Printf("firstBook: %+v\n", firstBook)
	lastBook := books[len(books)-1]
	fmt.Printf("lastBook: %+v\n", lastBook)

	book1 := Book{
		ID:     4,
		Author: "mark",
		Title:  "mark's book",
	}

	book1.SetAuthor("alan")
	fmt.Printf("book1: %+v\n", book1) // book1: {Id:4 Author:alan Title:mark's book}
}

func (b *Book) SetAuthor(s string) {
	b.Author = s
}
