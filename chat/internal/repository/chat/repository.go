package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lookandhate/microservice-courese/chat/internal/repository/model"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

func (r *PostgresRepository) Create(context.Context, *model.CreateChatModel) (*model.ChatModel, error) {
	panic("implement me")
}

func (r *PostgresRepository) CreateMessage(context.Context, *model.CreateMessageModel) (*model.MessageModel, error) {
	panic("Implement me")
}

func (r *PostgresRepository) Delete(context.Context, *model.DeleteChatModel) (bool, error) {
	panic("Implement me")
}

func (r *PostgresRepository) ChatExists(ctx context.Context, chatID int) (bool, error) {
	panic("Implement me")
}

func NewPostgresRepository(context context.Context, connectionDSN string) *PostgresRepository {
	pgx, err := pgxpool.New(context, connectionDSN)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &PostgresRepository{pgx: pgx}
}
