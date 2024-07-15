package user

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lookandhate/microservice-courese/auth/internal/model"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user *model.CreateUserRepositoryModel) (int, error) {
	builder := squirrel.Insert("users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("email", "password_hash", "name", "role").
		Values(user.Email, user.Password, user.Name, user.Role).
		Suffix("RETURNING id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	fmt.Println(sql, args)
	var id int
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *PostgresRepository) GetUser(ctx context.Context, id int) (user *model.UserModel, err error) {
	builder := squirrel.
		Select("id", "email", "password_hash", "name", "role").
		PlaceholderFormat(squirrel.Dollar).
		From("users").
		Where(squirrel.Eq{"id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	user = &model.UserModel{}

	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil

}
func (r *PostgresRepository) UpdateUser(ctx context.Context, user *model.UpdateUserModel) (updatedUser *model.UserModel, err error) {
	panic("implement me")
}
func (r *PostgresRepository) DeleteUser(ctx context.Context, id int) (err error) {
	panic("implement me")

}

func NewPostgresRepository(context context.Context, connectionDSN string) *PostgresRepository {
	pgx, err := pgxpool.Connect(context, connectionDSN)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &PostgresRepository{pgx: pgx}
}
