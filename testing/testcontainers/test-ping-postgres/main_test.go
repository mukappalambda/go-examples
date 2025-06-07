package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestPingPostgres(t *testing.T) {
	ctx := context.Background()
	dbName := "demo"
	dbUser := "postgres"
	dbPassword := "password"

	postgresImage := "postgres:17"
	reqOpts := []testcontainers.ContainerCustomizer{
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	}
	postgresContainer, err := postgres.Run(ctx, postgresImage, reqOpts...)
	if err != nil {
		t.Errorf("failed to run the postgres container: %s", err)
	}
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			fmt.Println(err)
		}
	}()
	connectionString, err := postgresContainer.ConnectionString(ctx, "")
	if err != nil {
		t.Errorf("failed to get connection string: %s", err)
	}
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		t.Errorf("failed to connect: %s", err)
	}
	defer conn.Close(context.Background())
}
