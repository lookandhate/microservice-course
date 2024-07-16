package repository

import (
	"context"

	"github.com/lookandhate/microservice-courese/auth/internal/model"
)

type UserRepository interface {
	CreateUser(context context.Context, user *model.CreateUserRepositoryModel) (int, error)
	GetUser(context context.Context, id int) (*model.UserModel, error)
	UpdateUser(context context.Context, updateUser *model.UpdateUserModel) (*model.UserModel, error)
	DeleteUser(context context.Context, id int) error
}
