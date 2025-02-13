package canvas_test

import (
	"sketch/internal/canvas"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawer_Draw(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
		requests []canvas.DrawRequest
	}{
		{
			name:     "1x1",
			expected: "*",
			requests: []canvas.DrawRequest{
				{Width: 1, Height: 1, Fill: "*"},
			},
		},
		{
			name:     "2x1",
			expected: "**",
			requests: []canvas.DrawRequest{
				{Width: 2, Height: 1, Fill: "*"},
			},
		},
		{
			name:     "1x2",
			expected: "*\n*",
			requests: []canvas.DrawRequest{
				{Width: 1, Height: 2, Fill: "*"},
			},
		},
		{
			name:     "2x2",
			expected: "**\n**",
			requests: []canvas.DrawRequest{
				{Width: 2, Height: 2, Fill: "*"},
			},
		},
		{
			name:     "1x1 in a different X axis",
			expected: " *",
			requests: []canvas.DrawRequest{
				{Width: 1, Height: 1, Fill: "*", X: 1},
			},
		},
		{
			name:     "1x1 in a different Y axis",
			expected: "\n*",
			requests: []canvas.DrawRequest{
				{Width: 1, Height: 1, Fill: "*", Y: 1},
			},
		},
		{
			name:     "1x1 (coordinates 0-0) + 1x1 (coordinates 2-0)",
			expected: "*+",
			requests: []canvas.DrawRequest{
				{
					Width:  1,
					Height: 1,
					Fill:   "*",
				},
				{
					X:      1,
					Width:  1,
					Height: 1,
					Fill:   "+",
				},
			},
		},
		{
			name:     "1x2 (coordinates 0-0) + 1x1 (coordinates 2-0)",
			expected: "* +\n*",
			requests: []canvas.DrawRequest{
				{
					Width:  1,
					Height: 2,
					Fill:   "*",
				},
				{
					X:      2,
					Width:  1,
					Height: 1,
					Fill:   "+",
				},
			},
		},
		{
			name:     "1x1 (coordinates 0-0) + 1x1 (coordinates 0-1)",
			expected: "*\n+",
			requests: []canvas.DrawRequest{
				{
					Width:  1,
					Height: 1,
					Fill:   "*",
				},
				{
					Width:  1,
					Height: 1,
					Fill:   "+",
					Y:      1,
				},
			},
		},
		{
			name: "x",
			expected: `              .......
              .......
              .......
              .......
              .......
              .......`,
			requests: []canvas.DrawRequest{
				{
					X:      14,
					Y:      0,
					Width:  7,
					Height: 6,
					Fill:   ".",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			drawer := canvas.NewDrawer()
			got, err := drawer.Draw(tc.requests)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestDrawer_DrawWithOutline(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
		requests []canvas.DrawRequest
	}{
		{
			name:     "3x3 with outline",
			expected: "  @@@\n  @+@\n  @@@",
			requests: []canvas.DrawRequest{
				{
					X:       2,
					Width:   3,
					Height:  3,
					Fill:    "+",
					Outline: "@",
				},
			},
		},
		{
			name: "...",
			expected: `


          XXXXXXXXXXXXXX
          XOOOOOOOOOOOOX
          XOOOOOOOOOOOOX
          XOOOOOOOOOOOOX
          XOOOOOOOOOOOOX
          XXXXXXXXXXXXXX`,
			requests: []canvas.DrawRequest{
				{
					Y:       3,
					X:       10,
					Width:   14,
					Height:  6,
					Fill:    "O",
					Outline: "X",
				},
			},
		},
		{
			name:     "outline character with 'none' fill",
			expected: "  XXXXX\n  X   X\n  XXXXX",
			requests: []canvas.DrawRequest{
				{X: 2, Width: 5, Height: 3, Fill: "none", Outline: "X"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			drawer := canvas.NewDrawer()
			got, err := drawer.Draw(tc.requests)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestDrawer_DrawMultiple(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
		requests []canvas.DrawRequest
	}{
		{
			name:     "should draw the intersection of two requests",
			expected: "🔥🔥🔥\n🔥🔥🔥\n💧💧💧",
			requests: []canvas.DrawRequest{
				{
					Width:  3,
					Height: 3,
					Fill:   "🔥",
				},
				{
					X:      0,
					Y:      2,
					Width:  3,
					Height: 1,
					Fill:   "💧",
				},
			},
		},
		{
			name: "Test fixture 1",
			expected: `

   @@@@@
   @XXX@  XXXXXXXXXXXXXX
   @@@@@  XOOOOOOOOOOOOX
          XOOOOOOOOOOOOX
          XOOOOOOOOOOOOX
          XOOOOOOOOOOOOX
          XXXXXXXXXXXXXX`,
			requests: []canvas.DrawRequest{
				{X: 3, Y: 2, Width: 5, Height: 3, Fill: "X", Outline: "@"},
				{X: 10, Y: 3, Width: 14, Height: 6, Fill: "O", Outline: "X"},
			},
		},
		{
			name: "Test fixture 2",
			expected: `              .......
              .......
              .......
OOOOOOOO      .......
O      O      .......
O    XXXXX    .......
OOOOOXXXXX
     XXXXX`,
			requests: []canvas.DrawRequest{
				{X: 14, Y: 0, Width: 7, Height: 6, Outline: "none", Fill: "."},
				{X: 0, Y: 3, Width: 8, Height: 4, Outline: "O", Fill: "none"},
				{X: 5, Y: 5, Width: 5, Height: 3, Outline: "X", Fill: "X"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			drawer := canvas.NewDrawer()
			got, err := drawer.Draw(tc.requests)
			expected := tc.expected

			assert.NoError(t, err)
			assert.Equal(t, expected, got)
		})
	}
}

func TestDrawer_Error(t *testing.T) {
	tests := []struct {
		name        string
		expectedErr error
	}{
		{
			name:        "when there are no requests, should return an error",
			expectedErr: canvas.ErrEmptyRequests,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			drawer := canvas.NewDrawer()
			_, err := drawer.Draw([]canvas.DrawRequest{})
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
