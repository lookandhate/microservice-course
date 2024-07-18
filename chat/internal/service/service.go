package service

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chatData *model.CreateChatRequest) (id int, err error)
	Delete(ctx context.Context, chatId int64) (err error)
	SendMessage(ctx context.Context, message *model.SendMessageRequest) (err error)
}
