package service

import (
	"context"

	"github.com/lookandhate/microservice-courese/auth/internal/model"
)

// GetUser validates user ID and after that tries to get user from repo
func (s *Service) GetUser(ctx context.Context, id int) (user *model.UserModel, err error) {
	// Check if ID is correct
	if err := s.validateID(id); err != nil {
		return nil, err
	}

	return s.repo.GetUser(ctx, id)
}
