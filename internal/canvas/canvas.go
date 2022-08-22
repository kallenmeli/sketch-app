package canvas

import (
	"time"

	"github.com/google/uuid"
)

const (
	EmptyChar = "none"
)

type Canvas struct {
	ID        string    `json:"id" db:"id"`
	Drawing   string    `json:"drawing" db:"drawing"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func NewCanvas(drawing string) Canvas {
	return Canvas{
		ID:        uuid.New().String(),
		Drawing:   drawing,
		CreatedAt: time.Now().UTC(),
	}
}
