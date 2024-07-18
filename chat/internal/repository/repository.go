package repository

import (
	"context"

	"github.com/lookandhate/microservice-courese/chat/internal/repository/model"
)

type ChatRepository interface {
	Create(context.Context, *model.CreateChatModel) (*model.ChatModel, error)
	CreateMessage(context.Context, *model.CreateMessageModel) (*model.MessageModel, error)
	Delete(context.Context, *model.DeleteChatModel) (bool, error)
}
