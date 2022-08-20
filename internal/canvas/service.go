package canvas

import (
	"context"
	"fmt"
)

type (
	service struct {
		repository Repository
	}
	Service interface {
		GetAll(ctx context.Context, id string) ([]Canvas, error)
		GetByID(ctx context.Context) (*Canvas, error)
		Save(ctx context.Context, canvas Canvas) error
	}
)

func NewService() Service {
	return nil
}

func (s service) GetAll(ctx context.Context) ([]Canvas, error) {
	drawings, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get drawings: %w", err)
	}
	return drawings, nil
}

func (s service) GetByID(ctx context.Context, id string) (*Canvas, error) {
	drawing, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get '%s': %w", id, err)
	}
	return &drawing, nil
}

func (s service) Save(ctx context.Context, canvas Canvas) error {
	//TODO implement me
	panic("implement me")
}
