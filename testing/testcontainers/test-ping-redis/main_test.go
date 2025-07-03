package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	tcredis "github.com/testcontainers/testcontainers-go/modules/redis"
)

func TestPostgres(t *testing.T) {
	ctx := context.Background()
	redisContainer, err := tcredis.Run(ctx, "redis:8")
	if err != nil {
		t.Errorf("failed to run redis container: %s", err.Error())
	}
	defer func() {
		if err := testcontainers.TerminateContainer(redisContainer); err != nil {
			t.Errorf("failed to terminate container: %s", err.Error())
		}
	}()
	connectionString, err := redisContainer.ConnectionString(ctx)
	if err != nil {
		t.Errorf("failed to get connection string: %s", err.Error())
	}
	urlOpt, err := redis.ParseURL(connectionString)
	if err != nil {
		t.Errorf("failed to parse url: %s", err)
	}
	client := redis.NewClient(urlOpt)
	defer client.Close()
	result, err := client.Ping(context.Background()).Result()
	if err != nil {
		t.Errorf("failed to ping redis server: %s", err)
	}
	fmt.Println("result:", result)
}
