package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func main() {
	dialer := &websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}
	ctx := context.Background()
	conn, _, err := dialer.DialContext(
		ctx,
		"ws://localhost:8000/s/test",
		http.Header{},
	)
	if err != nil {
		panic(err)
	}
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := conn.WriteMessage(
					websocket.TextMessage,
					[]byte("message from client1 @ "+time.Now().Format(time.RFC3339)),
				)
				if err != nil {
					fmt.Printf("write message error %v", err)
					return
				}
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Printf("receive: %s\n", msg)
	}
}
