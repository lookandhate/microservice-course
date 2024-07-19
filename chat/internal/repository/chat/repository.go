package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lookandhate/microservice-courese/chat/internal/repository/model"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

func (r *PostgresRepository) Create(ctx context.Context, request *model.CreateChatModel) (*model.ChatModel, error) {
	builder := squirrel.Insert("chats").
		PlaceholderFormat(squirrel.Dollar).
		Columns("created_at").
		Values(time.Now()).
		Suffix("returning id, created_at, updated_at")
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var chatModel model.ChatModel
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&chatModel.ID, &chatModel.CreatedAt, &chatModel.UpdatedAt)
	chatModel.UserIDs = request.UserIDs
	if err != nil {
		return nil, err
	}
	err = r.addUsersToChat(ctx, chatModel.ID, request.UserIDs)

	return &chatModel, err
}

// addUsersToChat creates chat members with given chatID and userIDSs
func (r *PostgresRepository) addUsersToChat(ctx context.Context, chatID int, userIDS []int64) error {
	// TODO COMPLETE ADDING USERS
	builder := squirrel.Insert("chat_members").PlaceholderFormat(squirrel.Dollar).Columns("user_id", "chat_id")
	for _, userID := range userIDS {
		builder = builder.Values(userID, chatID)
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	fmt.Println(sql, args, err)
	_, err = r.pgx.Exec(ctx, sql, args...)

	return err
}

// CreateMessage creates message
func (r *PostgresRepository) CreateMessage(context.Context, *model.CreateMessageModel) (*model.MessageModel, error) {
	panic("Implement me")
}

// Delete deletes chat
func (r *PostgresRepository) Delete(context.Context, *model.DeleteChatModel) (bool, error) {
	panic("Implement me")
}

// ChatExists checks if chat exists
func (r *PostgresRepository) ChatExists(ctx context.Context, chatID int) (bool, error) {
	panic("Implement me")
}

// NewPostgresRepository creates PostgresRepository instance
func NewPostgresRepository(context context.Context, connectionDSN string) *PostgresRepository {
	pgx, err := pgxpool.New(context, connectionDSN)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &PostgresRepository{pgx: pgx}
}
