package client

import "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/core/notes"

type services struct {
	noteStore notes.Store
}
