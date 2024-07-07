package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lookandhate/microservice-courese/auth/internal/grpc/auth"
	"github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/grpc"
)

const grpcPort = 50051

func main() {
	serverHost := fmt.Sprintf("localhost:%d", grpcPort) // Change host when use docker

	log.Printf("Serving at %v", serverHost)

	lis, err := net.Listen("tcp", serverHost)
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s := grpc.NewServer()
	auth_v1.RegisterAuthServer(s, &auth.Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC server %s", err)
	}
}
