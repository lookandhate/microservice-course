package service

import (
	"context"

	"github.com/lookandhate/microservice-courese/auth/internal/convertor"
	"github.com/lookandhate/microservice-courese/auth/internal/model"
	"github.com/lookandhate/microservice-courese/auth/internal/service"
)

// RegisterUser validates CreateUserModel, then passes it to repo layer and returns created user id
func (s *Service) RegisterUser(ctx context.Context, user *model.CreateUserModel) (id int, err error) {

	if user == nil {
		return 0, err
	}
	// Check user role has been passed correctly
	if user.Role == model.UserUnknownRole {
		err := service.ErrInvalidRole
		return 0, err
	}

	if user.PasswordConfirm != user.Password {
		return 0, service.ErrPasswordMismatch
	}

	createUserID, err := s.repo.CreateUser(ctx, convertor.CreateUserModelToRepo(user))
	if err != nil {
		return 0, err
	}
	return createUserID, nil

}
