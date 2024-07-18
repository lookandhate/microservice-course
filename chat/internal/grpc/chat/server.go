package chat

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit/v7"
	chatAPI "github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	chatAPI.UnimplementedChatServer
}

func (s *Server) Create(context context.Context, request *chatAPI.CreateRequest) (*chatAPI.CreateResponse, error) {
	log.Printf("Request: %#+v\n", request)

	return &chatAPI.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *Server) Delete(context context.Context, request *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Request: %#+v\n", request)

	return &emptypb.Empty{}, nil
}

func (s *Server) SendMessage(context context.Context, request *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Request: %#+v\n", request)

	return &emptypb.Empty{}, nil
}
