package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
)

var ErrInvalidAmount = errors.New("amount less than zero")

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	_ = ProcessPayment(ctx, logger, 10)
	_ = ProcessPayment(ctx, logger, -5)
}

func ProcessPayment(ctx context.Context, logger *slog.Logger, amount int) error {
	if amount < 0 {
		logger.With("main", "process-payment").ErrorContext(ctx, "payment failed", "amount", amount, "err", ErrInvalidAmount)
		return ErrInvalidAmount
	}

	logger.With("main", "process-payment").InfoContext(ctx, "payment succeeded", "amount", amount)
	return nil
}
