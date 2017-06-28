package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

const APIURL = "localhost:8080"

func main() {
	conn, err := net.Dial("tcp", APIURL)
	if err != nil {
		fmt.Printf("Error: connect to server: %s\n", err)
		os.Exit(1)
	}

	magic := "505249202a20485454502f322e300d0a0d0a534d0d0a0d0a"
	settings := "000000040000000000"
	headers := "0000550104000000018386458c62919aa5e1d7918c65a283ff418aa0e41d139d09b8f01e075f8b1d75d0620d263d4c4d65647a8d9acac8b4c7602bb6b81690bdff40027465864d833505b11f40899acac8b24d494f6a7f867df7dd6dd7ad"
	data := "00000c00010000000100000000070a05576f726c64"

	packets := []string{magic, settings, headers, data}

	for _, p := range packets {
		raw, err := hex.DecodeString(p)
		if err != nil {
			fmt.Printf("Error: cannot decode hex string: %s\n", err)
			os.Exit(1)
		}

		if _, err := conn.Write(raw); err != nil {
			fmt.Printf("Error: cannot decode hex string: %s\n", err)
			os.Exit(1)
		}
	}

	buf := make([]byte, 1024)

	fmt.Print("Info: received:\n")

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error: cannot read from socket: %s\n", err)
			os.Exit(1)
		}
		fmt.Print(string(buf[:n]))
	}
}
