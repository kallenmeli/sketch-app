package faker

import "sketch/internal/errors"

func NewError() error {
	return errors.Error("fake error")
}
