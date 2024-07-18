package repository

import "context"

type ChatRepository interface {
	CreateChat(ctx context.Context, chat *chat)
}
