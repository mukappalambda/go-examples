package store

import (
	"time"

	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/core/notes"
)

var FakeNotes = []notes.Note{
	{
		ID:        "1",
		Title:     "alpha's title",
		Content:   "alpha's content",
		CreatedAt: time.Now().Add(-2 * time.Second),
		UpdatedAt: time.Now().Add(-2 * time.Second),
	},
	{
		ID:        "2",
		Title:     "beta's title",
		Content:   "beta's content",
		CreatedAt: time.Now().Add(-1 * time.Second),
		UpdatedAt: time.Now().Add(-1 * time.Second),
	},
	{
		ID:        "3",
		Title:     "gamma's title",
		Content:   "gamma's content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
