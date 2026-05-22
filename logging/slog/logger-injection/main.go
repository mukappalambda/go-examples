package main

import (
	"context"
	"log/slog"
	"os"
)

type UserService struct {
	logger *slog.Logger
}

func (s *UserService) CreateUser(ctx context.Context, email string) error { //nolint:unparam
	// create user
	s.logger.With("component", "user_service").InfoContext(ctx, "create_user", "user_email", email)
	return nil
}

func main() {
	userService := &UserService{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
	ctx := context.Background()
	emails := []string{
		"alpha@example.com",
		"beta@example.com",
		"gamma@example.com",
	}
	for _, email := range emails {
		_ = userService.CreateUser(ctx, email)
	}
}
