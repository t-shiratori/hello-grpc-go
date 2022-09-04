package main

import (
	"context"
	"fmt"
	go_protocol_buffer "hello-grpc-go/go-protocol-buffer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	go_protocol_buffer.UnimplementedPinPonServiceServer
}

func (s *server) Send(ctx context.Context, req *go_protocol_buffer.PinPonRequest) (*go_protocol_buffer.PinPonResponse, error) {

	resWords := ""

	if req.Words == "Pin" {
		resWords = "Pon!"
	} else {
		resWords = "Please need words 'Pin'!"
	}

	res := &go_protocol_buffer.PinPonResponse{
		Words: resWords,
	}

	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	go_protocol_buffer.RegisterPinPonServiceServer(grpcServer, &server{})

	fmt.Println("server is runnig...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
