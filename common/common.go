package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// To convert Big Endian binary format of a 4 byte integer to int32
func FromBytes(b []byte) int32 {
	buf := bytes.NewReader(b)
	var result int32
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// To convert an int32 to a 4 byte Big Endian binary format
func ToBytes(i int32) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func WriteMsg(conn net.Conn, msg string) {
	// Send the size of the message to be sent
	conn.Write([]byte(ToBytes(int32(len([]byte(msg))))))
	// Send the message
	conn.Write([]byte(msg))
}

func ReadMsg(conn net.Conn) (string, error) {
	// Make a buffer to hold length of data
	lenBuf := make([]byte, 4)
	_, err := conn.Read(lenBuf)
	if err == io.EOF {
		fmt.Println("Connection Closed. Bye bye.")
		os.Exit(0)
	}
	if err != nil {
		return "", err
	}
	lenData := FromBytes(lenBuf)

	// Make a buffer to hold incoming data.
	buf := make([]byte, lenData)
	reqLen := 0
	// Keep reading data from the incoming connection into the buffer until all the data promised is
	// received
	for reqLen < int(lenData) {
		tempreqLen, err := conn.Read(buf[reqLen:])
		reqLen += tempreqLen
		if err == io.EOF {
			return "", fmt.Errorf("Received EOF before receiving all promised data.")
		}
		if err != nil {
			return "", fmt.Errorf("Error reading: %s", err.Error())
		}
	}
	return string(buf), nil
}
