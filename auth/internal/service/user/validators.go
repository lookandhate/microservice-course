package service

import "github.com/lookandhate/microservice-courese/auth/internal/service"

// validateID check passed ID and returns error if it is not correct
func (s *Service) validateID(id int) error {
	if id <= 0 {
		return service.ErrInvalidID
	}
	return nil
}
