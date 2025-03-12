package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/chat",
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create dial: %w", err)
	}
	defer conn.Close()
	fmt.Printf("guest@%s is online\n", conn.LocalAddr().String())
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read from stdin: %w", err)
		}
		data := strings.TrimSuffix(line, "\n")
		fmt.Printf("[server@%s]> %s\n", conn.RemoteAddr().String(), data)
		err = conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			return fmt.Errorf("failed to write message to server: %w", err)
		}
	}
}
