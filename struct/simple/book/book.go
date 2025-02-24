package book

type Book struct {
	id     int32
	author string
	title  string
}

func New(id int32, author, title string) *Book {
	return &Book{
		id:     id,
		author: author,
		title:  title,
	}
}

func (b *Book) ID() int32 {
	return b.id
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetTitle(title string) {
	b.title = title
}
