package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lookandhate/microservice-courese/auth/internal/grpc/auth"
	"github.com/lookandhate/microservice-courese/auth/internal/repository/user"
	service "github.com/lookandhate/microservice-courese/auth/internal/service/user"
	"github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/grpc"
)

const grpcPort = 50051

func main() {
	ctx := context.Background()
	serverHost := fmt.Sprintf("localhost:%d", grpcPort) // Change host when use docker

	log.Printf("Serving at %v", serverHost)

	userRepo := user.NewPostgresRepository(ctx, "host=localhost port=54320 dbname=auth user=POSTGRES_USER password=POSTGRES_PASSWORD sslmode=disable")
	userService := service.NewUserService(userRepo)
	server, err := auth.NewAuthServer(userService)
	if err != nil {
		log.Fatalf("failed to create auth server: %v", err)
	}

	lis, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s := grpc.NewServer()
	auth_v1.RegisterAuthServer(s, server)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
