package main

import (
	"fmt"
	"github.com/jonfk/golang-chat/tcp/common"
	"io"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

var (
	connections []net.Conn
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Save connection
		connections = append(connections, conn)
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	for {
		msg, err := common.ReadMsg(conn)
		if err != nil {
			if err == io.EOF {
				// Close the connection when you're done with it.
				removeConn(conn)
				conn.Close()
				return
			}
			log.Println(err)
			return
		}
		fmt.Printf("Message Received: %s\n", msg)
		broadcast(conn, msg)
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = range connections {
		if connections[i] == conn {
			break
		}
	}
	connections = append(connections[:i], connections[i+1:]...)
}

func broadcast(conn net.Conn, msg string) {
	for i := range connections {
		if connections[i] != conn {
			err := common.WriteMsg(connections[i], msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
