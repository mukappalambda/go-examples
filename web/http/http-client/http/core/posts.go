package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type Post struct {
	UserID int    `json:"userId,omitempty"`
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

func (p Post) Display() {
	fmt.Printf("%+v\n", p)
}

func PostFromReader(r io.Reader) (*Post, error) {
	var post Post
	if err := json.NewDecoder(r).Decode(&post); err != nil {
		return nil, err
	}
	return &post, nil
}

func PostsFromReader(r io.Reader) ([]*Post, error) {
	var posts []*Post
	if err := json.NewDecoder(r).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}

type Store interface {
	Get(ctx context.Context, id int) (*Post, error)
	Create(ctx context.Context, post *Post) (*Post, error)
	List(ctx context.Context) ([]*Post, error)
}
