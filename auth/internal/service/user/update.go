package service

import (
	"context"
	"errors"

	"github.com/lookandhate/microservice-courese/auth/internal/model"
	"github.com/lookandhate/microservice-courese/auth/internal/service"
)

// UpdateUser validates passed user data and updates user info
func (s *Service) UpdateUser(ctx context.Context, user *model.UpdateUserModel) (updatedUser *model.UserModel, err error) {
	if user == nil {
		err = errors.New("user is nil")
		return nil, err
	}

	if user.Role == int(model.UserUnknownRole) {
		err := service.ErrInvalidRole
		return nil, err
	}

	return s.repo.UpdateUser(ctx, user)
}
