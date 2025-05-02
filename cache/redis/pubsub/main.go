package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	redisOptions := &redis.Options{
		Addr:                  ":6379",
		DB:                    0,
		MaxRetries:            -1,
		DialTimeout:           10 * time.Second,
		ReadTimeout:           30 * time.Second,
		WriteTimeout:          30 * time.Second,
		ContextTimeoutEnabled: true,
	}
	client := redis.NewClient(redisOptions)
	channel := "mychannel"
	pubsub := client.Subscribe(context.Background(), channel)
	defer pubsub.Close()

	done := make(chan struct{})
	go func() {
		consumer(pubsub)
		close(done)
	}()
	msgs := []string{"go", "redis", "go-redis"}
	if err := sender(context.Background(), client, channel, 100*time.Millisecond, msgs...); err != nil {
		return err
	}
	<-done
	return nil
}

func sender(ctx context.Context, client *redis.Client, channel string, d time.Duration, msgs ...string) error {
	var errs []error
	for _, msg := range msgs {
		_, err := client.Publish(ctx, channel, msg).Result()
		if err != nil {
			errs = append(errs, err)
		}
		time.Sleep(d)
	}
	return errors.Join(errs...)
}

func consumer(pubsub *redis.PubSub) {
	for msg := range pubsub.Channel() {
		fmt.Printf("[Time]: %s [Payload]: %s\n", time.Now(), msg.Payload)
	}
}
