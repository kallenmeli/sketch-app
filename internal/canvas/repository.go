package canvas

import (
	"context"
	"database/sql"
	goerrors "errors"
	"fmt"
	"sketch/internal/errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound = errors.Error("not found")
)

type (
	Repository interface {
		GetByID(ctx context.Context, id string) (Canvas, error)
		Save(ctx context.Context, canvas Canvas) error
	}

	repository struct {
		db *sqlx.DB
	}
)

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, id string) (Canvas, error) {
	const query = "select id, drawing, created_at from drawings where id = $1"
	var canvas Canvas
	if err := r.db.GetContext(ctx, &canvas, query, id); err != nil {
		if goerrors.Is(err, sql.ErrNoRows) {
			return canvas, ErrNotFound
		}

		return canvas, fmt.Errorf("database err: %w", err)
	}
	return canvas, nil
}

func (r *repository) Save(ctx context.Context, canvas Canvas) error {
	const query = "insert into drawings (id, drawing, created_at) values (:id, :drawing, :created_at)"
	if _, err := r.db.NamedExecContext(ctx, query, canvas); err != nil {
		return fmt.Errorf("database err: %w", err)
	}
	return nil
}
