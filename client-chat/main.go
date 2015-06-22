package main

import (
	"bufio"
	"fmt"
	"github.com/jonfk/golang-chat/common"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to server through tcp.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go printOutput(conn)
	writeInput(conn)
}

func writeInput(conn *net.TCPConn) {
	fmt.Println("Enter text:")
	for {
		// Read from stdin.
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		common.WriteMsg(conn, text)
	}
}

func printOutput(conn *net.TCPConn) {
	for {

		msg, err := common.ReadMsg(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)

		// buf := make([]byte, 1024)
		// // Read the incoming connection into the buffer.
		// reqLen, err := conn.Read(buf)
		// if err != nil {
		// 	if err == io.EOF {
		// 		conn.Close()
		// 		fmt.Println("Connection was closed. Bye bye.")
		// 		os.Exit(0)
		// 	}
		// 	fmt.Println("Error reading:", err.Error())
		// }
		// fmt.Printf("%s\n", string(buf[:reqLen+1]))
	}
}
