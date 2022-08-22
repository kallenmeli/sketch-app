package faker

import (
	"sketch/internal/canvas"
	"testing"
)

func NewSingleDrawRequest(t *testing.T) canvas.DrawRequest {
	return NewDrawRequests(t)[0]
}

func NewDrawRequests(t *testing.T) canvas.DrawRequests {
	t.Helper()
	return canvas.DrawRequests{
		{
			X:       0,
			Y:       0,
			Width:   1,
			Height:  1,
			Outline: "@",
			Fill:    ".",
		},
	}
}

func NewInvalidDrawRequests(t *testing.T) canvas.DrawRequests {
	t.Helper()
	return canvas.DrawRequests{
		{
			X:       0,
			Y:       0,
			Width:   0,
			Height:  0,
			Outline: "@",
			Fill:    ".",
		},
		{
			X:       0,
			Y:       0,
			Width:   1,
			Height:  1,
			Outline: "@",
			Fill:    ".",
		},
	}
}
