package chat

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	chatAPI "github.com/lookandhate/microservice-courese/chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type Server struct {
	chatAPI.UnimplementedChatServer
}

func (s *Server) Create(context context.Context, request *chatAPI.CreateRequest) (*chatAPI.CreateResponse, error) {
	response := &chatAPI.CreateResponse{
		Id: gofakeit.Int64(),
	}
	log.Printf("Request: %#+v\nResponse: %#+v", request, response)
	return response, nil

}
func (s *Server) Delete(context context.Context, request *chatAPI.DeleteRequest) (*emptypb.Empty, error) {
	response := &emptypb.Empty{}
	log.Printf("Request: %#+v\nResponse: %#+v", request, response)
	return response, nil

}
func (s *Server) SendMessage(context context.Context, request *chatAPI.SendMessageRequest) (*emptypb.Empty, error) {
	response := &emptypb.Empty{}
	log.Printf("Request: %#+v\nResponse: %#+v", request, response)
	return response, nil
}
