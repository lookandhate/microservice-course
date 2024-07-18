package service

import (
	"github.com/lookandhate/microservice-courese/chat/internal/repository"
	"github.com/lookandhate/microservice-courese/chat/internal/service"
)

var _ service.ChatService = (*Service)(nil)

type Service struct {
	repo repository.ChatRepository
}

func NewService(repo repository.ChatRepository) *Service {
	return &Service{repo: repo}
}
