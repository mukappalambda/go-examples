package main

import (
	"context"
	"log/slog"
	"os"
)

const (
	KeyRequestID = "request_id"
	KeyUserID    = "user_id"
)

func main() {
	initialLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	requestID := "my-request-id"
	userID := 123
	logger := WithRequestLogger(initialLogger, requestID, userID)
	HandleRequest(context.Background(), logger)
}

func WithRequestLogger(logger *slog.Logger, requestID string, userID int) *slog.Logger {
	return logger.With(KeyRequestID, requestID, KeyUserID, userID)
}

func HandleRequest(ctx context.Context, logger *slog.Logger) {
	logger.InfoContext(ctx, "request started")
	logger.InfoContext(ctx, "request finished")
}
