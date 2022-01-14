package main

import "fmt"

type Book struct {
  Id int32
  Author string
  Title string
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
  {
    Id: 3,
    Author: "joe",
    Title: "joe's book",
  },
}

func main()  {
  for _, book := range books {
    fmt.Println(book.Id, book.Author, book.Title)
  }

  var oldBook Book
  for _, book := range books {
    if book.Author == "alex" {
      oldBook.Id = book.Id
      oldBook.Author = book.Author
      oldBook.Title = book.Title
    }
  }

  fmt.Println("oldBook:", oldBook)

  firstBook := books[0]
  fmt.Println(firstBook)
  lastBook := books[len(books) - 1]
  fmt.Println(lastBook)

  book1 := Book{
    Id: 4,
    Author: "mark",
    Title: "mark's book",
  }

  book1.SetAuthor("alan")
  fmt.Println(book1) // {4 alan mark's book}
  book1.SetAuthor2("mark")
  fmt.Println(book1) // {4 alan mark's book}
}

func (b *Book) SetAuthor(s string)  {
  b.Author = s
}

func (b Book) SetAuthor2(s string) {
  b.Author = s
}