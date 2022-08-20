package canvas

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		GetAll(ctx context.Context) ([]Canvas, error)
		GetByID(ctx context.Context, id string) (Canvas, error)
		Save(ctx context.Context, canvas Canvas) error
	}

	repository struct {
		db sqlx.DB
	}
)

func NewRepository(db sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]Canvas, error) {
	drawings := make([]Canvas, 0)
	const query = "select * from drawings"
	if err := r.db.GetContext(ctx, &drawings, query); err != nil {
		return nil, fmt.Errorf("failed to get all drawings from database: %w", err)
	}
	return drawings, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (Canvas, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Save(ctx context.Context, canvas Canvas) error {
	//TODO implement me
	panic("implement me")
}
