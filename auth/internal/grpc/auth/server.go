package auth

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit/v7"
	authAPI "github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	authAPI.UnimplementedAuthServer
}

func (s *Server) Get(context context.Context, request *authAPI.GetRequest) (*authAPI.GetResponse, error) {
	log.Printf("Request: %#+v\n", request)

	return &authAPI.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *Server) Create(context context.Context, request *authAPI.CreateRequest) (*authAPI.CreateResponse, error) {
	log.Printf("Request: %#+v\n", request)

	return &authAPI.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *Server) Update(context context.Context, request *authAPI.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Request: %#+v\n", request)

	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(context context.Context, request *authAPI.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Request: %#+v\n", request)

	return &emptypb.Empty{}, nil
}
