package repository

type Book struct {
	ID     int32  `form:"id" json:"id,omitempty"`
	Author string `form:"author" json:"author,omitempty"`
	Title  string `form:"title" json:"title,omitempty"`
}
