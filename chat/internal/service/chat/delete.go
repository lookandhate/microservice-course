package service

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/service"
)

func (s Service) Delete(ctx context.Context, chatID int) error {
	if err := s.validateID(ctx, chatID); err != nil {
		return err
	}
	if err := s.checkChatExists(ctx, chatID); err != nil {
		return service.ErrChatDoesNotExist
	}

	return nil
}
