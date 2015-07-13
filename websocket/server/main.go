package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
	log.Println("Serving")
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
