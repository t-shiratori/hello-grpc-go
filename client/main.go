package main

import (
	"context"
	"fmt"
	go_protocol_buffer "hello-grpc-go/go-protocol-buffer"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := go_protocol_buffer.NewPinPonServiceClient(conn)

	callPinPon(client)
}

func callPinPon(client go_protocol_buffer.PinPonServiceClient) {

	reqWords := ""

	if len(os.Args) >= 2 {
		reqWords = os.Args[1]
	}

	res, err := client.Send(context.Background(), &go_protocol_buffer.PinPonRequest{Words: reqWords})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetWords())
}
