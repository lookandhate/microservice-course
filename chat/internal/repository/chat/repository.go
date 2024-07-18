package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

func NewPostgresRepository(context context.Context, connectionDSN string) *PostgresRepository {
	pgx, err := pgxpool.New(context, connectionDSN)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &PostgresRepository{pgx: pgx}
}
