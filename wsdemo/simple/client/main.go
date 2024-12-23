package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/chat",
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("failed to create dial: %s\n", err)
	}
	defer conn.Close()
	fmt.Printf("guest@%s is online\n", conn.LocalAddr().String())
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read from stdin: %s\n", err)
			break
		}
		data := strings.TrimSuffix(line, "\n")
		fmt.Printf("[server@%s]> %s\n", conn.RemoteAddr().String(), data)
		err = conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Fatalf("failed to write message to server")
			break
		}
	}
}
