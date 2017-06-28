package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

const APIURL = "localhost:8080"

var conn net.Conn

func main() {
	var err error

	conn, err = net.Dial("tcp", APIURL)
	if err != nil {
		fmt.Printf("Error: connect to server: %s\n", err)
		os.Exit(1)
	}

	magic := decodeHex("505249202a20485454502f322e300d0a0d0a534d0d0a0d0a")
	send(magic)

	settings := decodeHex("000000040000000000")
	send(settings)

	headers := decodeHex("0000550104000000018386458c62919aa5e1d7918c65a283ff418aa0e41d139d09b8f01e075f8b1d75d0620d263d4c4d65647a8d9acac8b4c7602bb6b81690bdff40027465864d833505b11f40899acac8b24d494f6a7f867df7dd6dd7ad")
	send(headers)

	data := decodeHex("00000c00010000000100000000070a05576f726c64")
	send(data)

	listen()
}

func listen() {
	b := make([]byte, 1024)
	fmt.Print("Info: received:\n")
	for {
		n, err := conn.Read(b)
		if err != nil {
			fmt.Printf("Error: cannot read from socket: %s\n", err)
			os.Exit(1)
		}
		fmt.Print(string(b[:n]))
	}
}

func decodeHex(s string) []byte {
	raw, err := hex.DecodeString(s)
	if err != nil {
		fmt.Printf("Error: cannot decode hex string: %s\n", err)
		os.Exit(1)
	}

	return raw
}

func send(b []byte) {
	_, err := conn.Write(b)
	if err != nil {
		fmt.Printf("Error: cannot decode hex string: %s\n", err)
		os.Exit(1)
	}
}
