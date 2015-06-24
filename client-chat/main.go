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
	}
}
