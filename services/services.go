package services

import (
	"context"
	"fmt"

	"config"
)

// Service represents an interface for business operations.
type Service interface {
	DoSomething(ctx context.Context) error
}

// serviceImpl implements the Service interface.
type serviceImpl struct {
	cfg *config.Config
}

// NewService returns a new instance of the Service.
func NewService(cfg *config.Config) Service {
	return &serviceImpl{
		cfg: cfg,
	}
}

// DoSomething is an example method that demonstrates performing some business operation.
func (s *serviceImpl) DoSomething(ctx context.Context) error {
	// Implement your logic here, for example:
	fmt.Printf("Service running on port %s\n", s.cfg.Port)

	// Simulate a successful operation
	return nil
}
