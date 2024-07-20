package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lookandhate/microservice-courese/chat/internal/config"
	"github.com/lookandhate/microservice-courese/chat/internal/grpc/chat"
	repository "github.com/lookandhate/microservice-courese/chat/internal/repository/chat"
	"github.com/lookandhate/microservice-courese/chat/internal/service/chat"
	"github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.MustLoad()
	serverHost := fmt.Sprintf("localhost:%d", cfg.GPRC.Port) // Change host when use docker

	log.Printf("Serving at %v", serverHost)

	lis, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	ctx := context.Background()
	repo := repository.NewPostgresRepository(ctx, &cfg.Database)
	server := service.NewService(repo)

	s := grpc.NewServer()
	chatServer := chat.NewChatServer(server)
	chat_v1.RegisterChatServer(s, chatServer)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
