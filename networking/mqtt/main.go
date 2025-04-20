package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	topic    = flag.String("topic", "my-topic", "topic name")
	scheme   = flag.String("scheme", "tcp", "scheme")
	host     = flag.String("host", "localhost", "host")
	port     = flag.String("port", "1883", "port")
	clientId = flag.String("clientId", "my-client", "client id")
	payload  = flag.String("payload", "hello world", "payload")
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	var wg sync.WaitGroup

	server := fmt.Sprintf("%s://%s:%s", *scheme, *host, *port)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server).SetClientID(*clientId)
	opts.SetOnConnectHandler(onConn)
	wg.Add(1)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		defer wg.Done()
		fmt.Printf("%v - Topic: %s; payload: %s", time.Now().Truncate(time.Microsecond), msg.Topic(), msg.Payload())
	})
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("%s", token.Error().Error())
	}

	if token := client.Subscribe(*topic, 0, nil); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe topic: %s", token.Error().Error())
	}

	wg.Add(1)
	var e error
	go func() {
		defer wg.Done()
		if token := client.Publish(*topic, 0, false, *payload); token.Wait() && token.Error() != nil {
			e = token.Error()
		}
	}()
	wg.Wait()
	if e != nil {
		return e
	}
	client.Disconnect(250)
	return nil
}

func onConn(client mqtt.Client) {
	fmt.Printf("%v - Client connected\n", time.Now().Truncate(time.Microsecond))
}
