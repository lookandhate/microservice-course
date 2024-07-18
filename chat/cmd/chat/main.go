package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lookandhate/microservice-courese/chat/internal/grpc/chat"
	repository "github.com/lookandhate/microservice-courese/chat/internal/repository/chat"
	"github.com/lookandhate/microservice-courese/chat/internal/service/chat"
	"github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
	"google.golang.org/grpc"
)

const grpcPort = 50052

func main() {
	// TODO: ADD CONFIG
	serverHost := fmt.Sprintf("localhost:%d", grpcPort) // Change host when use docker

	log.Printf("Serving at %v", serverHost)

	lis, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	ctx := context.Background()
	repo := repository.NewPostgresRepository(ctx, "host=localhost port=54321 dbname=chat user=POSTGRES_USER password=POSTGRES_PASSWORD sslmode=disable")
	server := service.NewService(repo)

	s := grpc.NewServer()
	chatServer := chat.NewChatServer(server)
	chat_v1.RegisterChatServer(s, chatServer)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
