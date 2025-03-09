package notes

import (
	"context"
	"time"
)

type Note struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Store interface {
	Get(ctx context.Context, id string) (Note, error)
	Create(ctx context.Context, note Note) (Note, error)
	List(ctx context.Context, filters ...string) ([]Note, error)
	Update(ctx context.Context, note Note, fieldpaths ...string) (Note, error)
	Delete(ctx context.Context, id string) error
}
