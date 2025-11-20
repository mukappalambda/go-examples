package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		logger = flag.String("logger", "", "logger adapter")
		debug  = flag.Bool("debug", false, "debug mode")
		trace  = flag.Bool("trace", false, "trace mode")
	)
	flag.Parse()
	if len(*logger) == 0 {
		*logger = "std"
	}

	loggers := map[string]watermill.LoggerAdapter{
		"std":  watermill.NewStdLoggerWithOut(os.Stdout, *debug, *trace),
		"slog": watermill.NewSlogLogger(slog.Default()),
	}
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{BlockPublishUntilSubscriberAck: true},
		loggers[*logger],
	)
	topicName := "example.topic"
	var e error
	go func() {
		defer func() {
			err := pubSub.Close()
			if err != nil {
				e = fmt.Errorf("failed to close channel: %s", err)
			}
		}()
		msgs := []*message.Message{
			message.NewMessage(watermill.NewUUID(), []byte("Do the laundry")),
			message.NewMessage(watermill.NewUUID(), []byte("Go to workout")),
			message.NewMessage(watermill.NewUUID(), []byte("Prepare the breakfast")),
		}
		err := pubSub.Publish(topicName, msgs...)
		if err != nil {
			_ = pubSub.Close()
			e = fmt.Errorf("failed to publish messages: %s", err)
			return
		}
	}()
	if e != nil {
		return e
	}
	msgs, err := pubSub.Subscribe(context.Background(), topicName)
	if err != nil {
		return fmt.Errorf("failed to subscribe messages: %s", err)
	}
	for msg := range msgs {
		fmt.Printf("[uuid]: %s [payload]: %s\n", msg.UUID, string(msg.Payload))
		msg.Ack()
	}
	return nil
}
