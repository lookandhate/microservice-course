package user

import (
	"context"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	repository "github.com/lookandhate/microservice-courese/auth/internal/repository/model"
	"github.com/lookandhate/microservice-courese/auth/internal/service/model"
)

type PostgresRepository struct {
	pgx *pgxpool.Pool
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user *repository.CreateUserModel) (int, error) {
	builder := squirrel.Insert("users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("email", "password_hash", "name", "role").
		Values(user.Email, user.Password, user.Name, user.Role).
		Suffix("RETURNING id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *PostgresRepository) GetUser(ctx context.Context, id int) (*model.UserModel, error) {
	builder := squirrel.
		Select("id", "email", "password_hash", "name", "role").
		PlaceholderFormat(squirrel.Dollar).
		From("users").
		Where(squirrel.Eq{"id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user model.UserModel

	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, user *model.UpdateUserModel) (*model.UserModel, error) {
	builder := squirrel.Update("users").PlaceholderFormat(squirrel.Dollar).Where(squirrel.Eq{"id": user.ID})

	if user.Password != nil {
		builder = builder.Set("password_hash", user.Password)
	}
	if user.Name != nil {
		builder = builder.Set("name", user.Name)
	}
	if user.Role != int(model.UserUnknownRole) {
		builder = builder.Set("role", user.Role)
	}
	if user.Email != nil {
		builder = builder.Set("email", user.Email)
	}

	builder = builder.Set("updated_at", time.Now()).Suffix("RETURNING id, email, name, role, created_at, updated_at")
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	updatedUser := &model.UserModel{}
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(
		&updatedUser.ID,
		&updatedUser.Email,
		&updatedUser.Name,
		&updatedUser.Role,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id int) error {
	builder := squirrel.Delete("users").Where(squirrel.Eq{"id": id})
	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	log.Println(sql, args)
	_, err = r.pgx.Exec(ctx, sql, args...)
	return err
}

func (r *PostgresRepository) CheckUserExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	// using Prefix and suffix for EXIST query
	builder := squirrel.Select("").
		PlaceholderFormat(squirrel.Dollar).
		Prefix("SELECT EXISTS (").
		From("users").
		Where(squirrel.Eq{"id": id}).
		Suffix(")")
	sql, args, err := builder.ToSql()

	if err != nil {
		return false, err
	}
	err = r.pgx.QueryRow(ctx, sql, args...).Scan(&exists)
	return exists, err
}

func NewPostgresRepository(context context.Context, connectionDSN string) *PostgresRepository {
	pgx, err := pgxpool.New(context, connectionDSN)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &PostgresRepository{pgx: pgx}
}
