package main

import (
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Info("user action", "user_id", 42, "action", "login", "success", true, "latency_ms", 137)
}
