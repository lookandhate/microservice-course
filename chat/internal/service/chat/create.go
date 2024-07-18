package chat

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/model"
)

func (s Service) Create(ctx context.Context, chat *model.CreateChatRequest) (int, error) {
	panic("implement me")
}

func (s Service) SendMessage(ctx context.Context, message *model.SendMessageRequest) error {
	panic("implement me")
}
