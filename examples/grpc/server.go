package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const Port = 8081

type apiServer struct{}

func (s *apiServer) Hello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	fmt.Printf("Info: request with name %s\n", in.Name)

	res := &HelloResponse{
		Message: fmt.Sprint("Hello ", in.Name),
	}

	return res, nil
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprint(":", Port))
	if err != nil {
		fmt.Printf("Error: cannot listen on port %d: %s\n", err)
		os.Exit(1)
	}

	server := grpc.NewServer()

	RegisterAPIServer(server, &apiServer{})

	fmt.Printf("Listening on port: %d\n", Port)
	if err := server.Serve(l); err != nil {
		fmt.Printf("Error: cannot start grpc server: %s\n", err)
		os.Exit(1)
	}

}
