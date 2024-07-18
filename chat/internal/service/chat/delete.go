package chat

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/service"
)

var _ service.ChatService = (*Service)(nil)

func (s Service) Delete(ctx context.Context, chatID int64) error {
	panic("implement me")
}
