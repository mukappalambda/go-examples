package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

type MyHandler struct {
	w io.Writer
}

func newMyHandler(w io.Writer) *MyHandler {
	return &MyHandler{w: w}
}

func (m *MyHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (m *MyHandler) Handle(ctx context.Context, r slog.Record) error {
	return slog.NewJSONHandler(m.w, nil).Handle(ctx, r)
}

func (m *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return slog.Default().Handler()
}

func (m *MyHandler) WithGroup(name string) slog.Handler {
	return slog.Default().Handler()
}

var _ slog.Handler = (*MyHandler)(nil)

func main() {
	logFile, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Different ways to create the *slog.Logger object
	loggers := map[string]*slog.Logger{
		"default-handler-v1":        slog.Default(),
		"default-handler-v2":        slog.New(slog.Default().Handler()),
		"stdout-json-handler":       slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		"stdout-debug-json-handler": slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
		"file-json-handler":         slog.New(slog.NewJSONHandler(logFile, nil)),
		"stdout-text-handler":       slog.New(slog.NewTextHandler(os.Stdout, nil)),
		"stderr-custom-handler":     slog.New(newMyHandler(os.Stderr)),
	}

	for name, logger := range loggers {
		fmt.Println(name)
		process(logger)
		fmt.Println()
	}
}

func process(logger *slog.Logger) {
	logger.Debug("debug msg")
	logger.Info("info msg")
	logger.Warn("warn msg")
	logger.Error("error msg")
}
