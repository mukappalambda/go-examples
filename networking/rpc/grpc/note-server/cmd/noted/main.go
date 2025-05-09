package main

import (
	"fmt"
	"os"

	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/cmd/noted/command"
)

func main() {
	app := command.App()
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "note-server: %s\n", err)
		os.Exit(1)
	}
}
