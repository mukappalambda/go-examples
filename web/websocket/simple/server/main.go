package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/chat", HandleWS())
	server := &http.Server{Addr: ":8080", ReadTimeout: time.Second}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func HandleWS() http.HandlerFunc {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "failed to upgrade to the Websocket connection", http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		log.Printf("guest@%s logs in\n", conn.RemoteAddr().String())
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Printf("failed to read message: %s\n", err)
				http.Error(w, "failed to read message temporarily: %s\n", http.StatusInternalServerError)
				return
			}
			fmt.Printf("[guest@%s]> %s\n", conn.RemoteAddr().String(), string(p))
			data := []byte("you said: ")
			data = append(data, p...)
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Printf("failed to write message: %s\n", err)
			}
		}
	}
}
