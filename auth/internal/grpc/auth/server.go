package auth

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	authapi "github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type ServerAPI struct {
	authapi.UnimplementedAuthServer
}

func (s *ServerAPI) Get(context context.Context, request *authapi.GetRequest) (*authapi.GetResponse, error) {
	response := &authapi.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}
	log.Printf("Request: %#+v\nResponse: %#+v", request, response)

	return response, nil

}

func (s *ServerAPI) Create(context context.Context, request *authapi.CreateRequest) (*authapi.CreateResponse, error) {
	response := &authapi.CreateResponse{Id: gofakeit.Int64()}
	log.Printf("Request: %#+v\nResponse: %#+v", request, response)

	return response, nil

}

func (s *ServerAPI) Update(context context.Context, request *authapi.UpdateRequest) (*emptypb.Empty, error) {
	response := &emptypb.Empty{}
	log.Printf("Request: %#+v\nResponse: %#+v", request, response)
	return response, nil

}
func (s *ServerAPI) Delete(context context.Context, request *authapi.DeleteRequest) (*emptypb.Empty, error) {
	response := &emptypb.Empty{}
	log.Printf("Request: %#+v\nResponse: %#+v", request, response)
	return response, nil

}
