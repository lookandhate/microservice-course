package service

import (
	"context"
)

// DeleteUser deletes user by given ID if it is correct.
func (s *Service) DeleteUser(ctx context.Context, id int) error {
	// Check if ID is correct
	if err := s.validateID(id); err != nil {
		return err
	}

	if err := s.checkUserExists(ctx, id); err != nil {
		return err
	}

	return s.repo.DeleteUser(ctx, id)
}
