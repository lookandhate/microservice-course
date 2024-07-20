package repository

import (
	"context"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lookandhate/microservice-courese/chat/internal/config"
	"github.com/lookandhate/microservice-courese/chat/internal/repository/model"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

// CreateChat creates chat with chat members.
func (r *PostgresRepository) CreateChat(ctx context.Context, request *model.CreateChatModel) (*model.ChatModel, error) {
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

// addUsersToChat creates chat members with given chatID and userIDSs.
func (r *PostgresRepository) addUsersToChat(ctx context.Context, chatID int, userIDs []int64) error {
	builder := squirrel.
		Insert("chat_members").
		PlaceholderFormat(squirrel.Dollar).
		Columns("user_id", "chat_id")

	for _, userID := range userIDs {
		builder = builder.Values(userID, chatID)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pgx.Exec(ctx, sql, args...)

	return err
}

// CreateMessage creates message.
func (r *PostgresRepository) CreateMessage(ctx context.Context, message *model.CreateMessageModel) (*model.MessageModel, error) {
	builder := squirrel.Insert("message").
		PlaceholderFormat(squirrel.Dollar).
		Columns("author_id", "content", "chat_id").
		Values(message.AuthorID, message.Content, message.ChatID).
		Suffix("returning id, author_id, content, chat_id, created_at, updated_at")

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var createdMessage model.MessageModel
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(
		&createdMessage.ID,
		&createdMessage.Author,
		&createdMessage.Content,
		&createdMessage.ChatID,
		&createdMessage.CreatedAt,
		&createdMessage.UpdatedAt)

	return &createdMessage, err
}

// Delete deletes chat.
func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
	builder := squirrel.Delete("chats").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.pgx.Exec(ctx, sql, args...)
	return err
}

// ChatExists checks if chat exists.
func (r *PostgresRepository) ChatExists(ctx context.Context, chatID int) (bool, error) {
	var exists bool
	// using Prefix and suffix for EXIST query
	builder := squirrel.Select("").
		PlaceholderFormat(squirrel.Dollar).
		Prefix("SELECT EXISTS (").
		From("chats").
		Where(squirrel.Eq{"id": chatID}).
		Suffix(")")
	sql, args, err := builder.ToSql()

	if err != nil {
		return false, err
	}
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&exists)
	return exists, err
}

// NewPostgresRepository creates PostgresRepository instance.
func NewPostgresRepository(context context.Context, dbConfig config.DatabaseConfig) *PostgresRepository {
	pgx, err := pgxpool.New(context, dbConfig.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &PostgresRepository{pgx: pgx}
}
