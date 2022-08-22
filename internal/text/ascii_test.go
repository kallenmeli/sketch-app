package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASCIIChar_Validate(t *testing.T) {
	tests := []struct {
		name   string
		a      ASCIIChar
		assert func(t *testing.T, err error)
	}{
		{
			name: "when empty, should return nil",
			a:    "",
			assert: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:   "when length is greater than 1, should return error",
			a:      "aa",
			assert: func(t *testing.T, err error) { assert.ErrorIs(t, err, ErrInvalidASCIIChar) },
		},
		{
			name:   "when value is greater than unicode.MaxASCII, should return error",
			a:      "😆",
			assert: func(t *testing.T, err error) { assert.ErrorIs(t, err, ErrInvalidASCIIChar) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.a.Validate()
			tt.assert(t, err)
		})
	}
}
