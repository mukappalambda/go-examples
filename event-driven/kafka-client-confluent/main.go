package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var brokers = flag.String("brokers", "", "brokers")

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "[kafka-client]: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	if len(*brokers) == 0 {
		flag.PrintDefaults()
		return fmt.Errorf("missing brokers")
	}
	conf := &kafka.ConfigMap{
		"bootstrap.servers": *brokers,
		"group.id":          "test",
	}
	c, err := kafka.NewConsumer(conf)
	if err != nil {
		return fmt.Errorf("failed to new consumer: %s", err)
	}
	defer c.Close()
	cgm, err := c.GetConsumerGroupMetadata()
	if err != nil {
		return fmt.Errorf("failed to get metadata: %s", err)
	}
	fmt.Printf("consumer group metadata: %s\n", cgm)
	return nil
}
