package main

import (
	"fmt"
	"log"
	"net"

	"grpc/chat"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")
	fmt.Println(fmt.Sprintf("abc:%d", 9000))
	lis, err := net.Listen("tcp", fmt.Sprintf("abc:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
