package canvas

import (
	"context"
	"fmt"
)

type (
	service struct {
		repository Repository
		drawer     Drawer
	}
	Service interface {
		GetByID(ctx context.Context, id string) (*Canvas, error)
		Save(ctx context.Context, requests DrawRequests) (*DrawResponse, error)
	}
)

func NewService(repository Repository, drawer Drawer) Service {
	return &service{
		repository: repository,
		drawer:     drawer,
	}
}

func (s service) GetByID(ctx context.Context, id string) (*Canvas, error) {
	drawing, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get '%s': %w", id, err)
	}
	return &drawing, nil
}

func (s service) Save(ctx context.Context, request DrawRequests) (*DrawResponse, error) {
	draw, err := s.drawer.Draw(request)
	if err != nil {
		return nil, fmt.Errorf("fail to draw: %w", err)
	}

	canvas := NewCanvas(draw)
	if err := s.repository.Save(ctx, canvas); err != nil {
		return nil, fmt.Errorf("error saving canvas: %w", err)
	}

	return &DrawResponse{
		ID:      canvas.ID,
		Drawing: draw,
	}, nil
}
