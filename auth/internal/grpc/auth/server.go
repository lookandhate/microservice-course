package auth

import (
	"context"
	"log"

	"github.com/lookandhate/microservice-courese/auth/internal/convertor"
	"github.com/lookandhate/microservice-courese/auth/internal/service"
	authAPI "github.com/lookandhate/microservice-courese/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	authAPI.UnimplementedAuthServer
	userService service.UserService
}

func (s *Server) Get(ctx context.Context, request *authAPI.GetRequest) (*authAPI.GetResponse, error) {
	user, err := s.userService.GetUser(ctx, int(request.Id))
	if err != nil {
		return nil, err
	}

	return convertor.UserModelToGetResponseProto(user), err
}

func (s *Server) Create(ctx context.Context, request *authAPI.CreateRequest) (*authAPI.CreateResponse, error) {
	userID, err := s.userService.RegisterUser(ctx, convertor.CreateUserFromProto(request))
	if err != nil {
		return nil, err
	}

	return &authAPI.CreateResponse{Id: int64(userID)}, nil
}

func (s *Server) Update(ctx context.Context, request *authAPI.UpdateRequest) (*emptypb.Empty, error) {
	_, err := s.userService.UpdateUser(ctx, convertor.UserUpdateFromProto(request))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(context context.Context, request *authAPI.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Request: %#+v\n", request)

	return &emptypb.Empty{}, nil
}

func NewAuthServer(service service.UserService) (*Server, error) {
	return &Server{
		userService: service,
	}, nil
}
