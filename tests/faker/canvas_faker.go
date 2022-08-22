package faker

import (
	"sketch/internal/canvas"
	"testing"
)

func NewCanvas(t *testing.T) canvas.Canvas {
	t.Helper()
	return canvas.NewCanvas(":)")
}
