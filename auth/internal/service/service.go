package service

import (
	"context"

	"github.com/lookandhate/microservice-courese/auth/internal/service/model"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *model.CreateUserModel) (id int, err error)
	GetUser(ctx context.Context, id int) (user *model.UserModel, err error)
	UpdateUser(ctx context.Context, user *model.UpdateUserModel) (updatedUser *model.UserModel, err error)
	DeleteUser(ctx context.Context, id int) (err error)
}
