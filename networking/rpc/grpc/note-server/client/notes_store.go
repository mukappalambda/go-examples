package client

import (
	"context"
	"fmt"

	notesapi "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"
	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/core/notes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type remoteNotes struct {
	client notesapi.NoteServiceClient
}

var _ notes.Store = (*remoteNotes)(nil)

func NewRemoteNoteStore(client notesapi.NoteServiceClient) notes.Store {
	return &remoteNotes{
		client: client,
	}
}

func (r *remoteNotes) Get(ctx context.Context, id string) (notes.Note, error) {
	resp, err := r.client.GetNote(ctx, &notesapi.GetNoteRequest{
		Id: id,
	})
	if err != nil {
		return notes.Note{}, fmt.Errorf("failed to get note: %w", err)
	}
	return notes.Note{
		ID:        resp.Note.Id,
		Title:     resp.Note.Title,
		Content:   resp.Note.Content,
		CreatedAt: resp.Note.CreatedAt.AsTime(),
		UpdatedAt: resp.Note.UpdatedAt.AsTime(),
	}, nil
}

func (r *remoteNotes) Create(ctx context.Context, note notes.Note) (notes.Note, error) {
	n := &notesapi.Note{
		Id:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: timestamppb.New(note.UpdatedAt),
	}
	resp, err := r.client.CreateNote(ctx, &notesapi.CreateNoteRequest{
		Note: n,
	})
	if err != nil {
		return notes.Note{}, err
	}
	return notes.Note{
		ID:        resp.Note.Id,
		Title:     resp.Note.Title,
		Content:   resp.Note.Content,
		CreatedAt: resp.Note.CreatedAt.AsTime(),
		UpdatedAt: resp.Note.UpdatedAt.AsTime(),
	}, nil
}

func (r *remoteNotes) List(ctx context.Context, filters ...string) ([]notes.Note, error) {
	resp, err := r.client.List(ctx, &notesapi.ListNotesRequest{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	var noteList []notes.Note
	for _, n := range resp.GetNotes() {
		newN := notes.Note{
			ID:        n.Id,
			Title:     n.Title,
			Content:   n.Content,
			CreatedAt: n.CreatedAt.AsTime(),
			UpdatedAt: n.UpdatedAt.AsTime(),
		}
		noteList = append(noteList, newN)
	}
	return noteList, nil
}

func (r *remoteNotes) Update(ctx context.Context, note notes.Note, fieldpaths ...string) (notes.Note, error) {
	panic("unimplemented")
}

func (r *remoteNotes) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}
