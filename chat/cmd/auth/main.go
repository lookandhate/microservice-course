package main

import (
	"fmt"
	"github.com/lookandhate/microservice-courese/chat/internal/grpc/chat"
	"github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

const grpcPort = 50052

func main() {
	serverHost := fmt.Sprintf("localhost:%d", grpcPort) // Change host when use docker

	log.Printf("Serving at %v", serverHost)

	lis, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s := grpc.NewServer()
	chat_v1.RegisterChatServer(s, &chat.Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
