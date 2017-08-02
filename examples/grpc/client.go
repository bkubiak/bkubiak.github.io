package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	APIURL = "localhost:8080"
	Name   = "World"
)

func main() {
	conn, err := grpc.Dial(APIURL, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		fmt.Printf("Error: connect to server: %s\n", err)
		os.Exit(1)
	}

	client := NewAPIClient(conn)

	payload := &HelloRequest{
		Name: Name,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Hello(ctx, payload)
	if err != nil {
		fmt.Printf("Error: cannot send request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Info: response: %s\n", res.Message)
}
