package library

import (
	"fmt"

	"github.com/mukappalambda/go-examples/struct/simple/book"
)

type Library struct {
	books []*book.Book
}

func New(books []*book.Book) *Library {
	return &Library{books: books}
}

func (l *Library) AddBooks(books ...*book.Book) {
	for _, book := range books {
		l.addBook(book)
	}
}

func (l *Library) addBook(book *book.Book) {
	l.books = append(l.books, book)
}

func (l *Library) DisplayBooks() {
	for _, book := range l.books {
		fmt.Printf("[ID]: %d [Author]: %q [Title]: %q\n", book.ID(), book.Author(), book.Title())
	}
}
