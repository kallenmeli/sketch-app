package canvas_test

import (
	"sketch/internal/canvas"
	"sketch/internal/text"
	"sketch/tests/faker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawRequest_GetOutlineChar(t *testing.T) {
	tests := []struct {
		name     string
		char     text.ASCIIChar
		expected string
	}{
		{
			name:     "when empty, should return empty string",
			char:     "",
			expected: "",
		},
		{
			name:     "when filled with a single character, should return that character",
			char:     "a",
			expected: "a",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			request := faker.NewSingleDrawRequest(t)
			request.Outline = tc.char
			char := request.GetOutlineChar()
			assert.EqualValues(t, tc.expected, char)
		})
	}
}

func TestDrawRequest_Validate(t *testing.T) {
	type fields struct {
		X       int
		Y       int
		Width   int
		Height  int
		Outline text.ASCIIChar
		Fill    text.ASCIIChar
	}
	tests := []struct {
		name   string
		fields fields
		assert func(t *testing.T, err error)
	}{
		{
			name:   "when fill and outline are empty, should return an error",
			fields: fields{},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "at least one")
			},
		},
		{
			name: "when fill is an invalid ascii character, should return an error",
			fields: fields{
				Fill: "😥",
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, text.ErrInvalidASCIIChar)
			},
		},
		{
			name: "when outline is an invalid ascii character, should return an error",
			fields: fields{
				Outline: "😥",
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, text.ErrInvalidASCIIChar)
			},
		},
		{
			name: "when width is less than 1, should return an error",
			fields: fields{
				Width: 0,
				Fill:  "*",
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "width")
			},
		},
		{
			name: "when height is less than 1, should return an error",
			fields: fields{
				Height: 0,
				Fill:   "*",
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "height")
			},
		},
		{
			name: "when x is less than 0, should return an error",
			fields: fields{
				X:    -1,
				Fill: "*",
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "coordinates must be equal or greater than zero")
			},
		},
		{
			name: "when y is less than 0, should return an error",
			fields: fields{
				Y:    -1,
				Fill: "*",
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "coordinates must be equal or greater than zero")
			},
		},
		{
			name: "when all fields are valid, should return no error",
			fields: fields{
				X:       1,
				Y:       2,
				Width:   3,
				Height:  4,
				Outline: "a",
				Fill:    "b",
			},
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := canvas.DrawRequest{
				X:       tt.fields.X,
				Y:       tt.fields.Y,
				Width:   tt.fields.Width,
				Height:  tt.fields.Height,
				Outline: tt.fields.Outline,
				Fill:    tt.fields.Fill,
			}
			err := request.Validate()
			tt.assert(t, err)
		})
	}
}

func TestDrawRequests_Validate(t *testing.T) {
	tests := []struct {
		name     string
		requests canvas.DrawRequests
		assert   func(t *testing.T, err error)
	}{
		{
			name:     "when there are no requests, should return an error",
			requests: canvas.DrawRequests{},
			assert: func(t *testing.T, err error) {
				assert.ErrorIs(t, canvas.ErrEmptyRequests, err)
			},
		},
		{
			name:     "when theres is one invalid requests, should return an error",
			requests: faker.NewInvalidDrawRequests(t),
			assert: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name:     "when there are no errors in the requests, should return nil",
			requests: faker.NewDrawRequests(t),
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.requests.Validate()
			tc.assert(t, err)
		})
	}
}
