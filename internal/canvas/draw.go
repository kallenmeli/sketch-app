package canvas

import (
	"sketch/internal/errors"
	"sketch/internal/text"
	"strings"
)

type (
	Draw [][]string

	DrawRequest struct {
		X       int            `json:"x" validate:"required"`
		Y       int            `json:"y" validate:"required"`
		Width   int            `json:"width" validate:"required"`
		Height  int            `json:"height" validate:"required"`
		Outline text.ASCIIChar `json:"outline"`
		Fill    text.ASCIIChar `json:"fill"`
	}

	DrawRequests []DrawRequest

	DrawResponse struct {
		ID      string `json:"id"`
		Drawing string `json:"canvas"`
	}
)

func (d Draw) String() string {
	height := len(d)
	result := strings.Builder{}
	for i, row := range d {
		for _, value := range row {
			result.WriteString(value)
		}
		isFinalRow := i == height-1
		if !isFinalRow {
			result.WriteString("\n")
		}
	}
	return result.String()
}

var (
	ErrEmptyRequests = errors.Error("at least one request is required")
)

func NewDraw(width, height int) Draw {
	draw := make(Draw, height)
	for i := 0; i < height; i++ {
		draw[i] = make([]string, width)
	}
	return draw
}

func (d DrawRequests) Validate() error {
	if len(d) == 0 {
		return ErrEmptyRequests
	}

	for _, request := range d {
		if err := request.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (d DrawRequest) GetFillChar() string {
	if d.Fill == EmptyChar {
		return " "
	}
	return string(d.Fill)
}

func (d DrawRequest) GetOutlineChar() string {
	if d.Outline == EmptyChar {
		return ""
	}
	return string(d.Outline)
}

func (d DrawRequest) WidthEnd() int {
	return d.X + d.Width
}

func (d DrawRequest) HeightEnd() int {
	return d.Y + d.Height
}

func (d DrawRequest) IsLastRow(row int) bool {
	return row == d.HeightEnd()-1
}

func (d DrawRequest) IsFirstRow(row int) bool {
	return row == d.Y
}

func (d DrawRequest) Validate() error {
	isEmpty := func(value text.ASCIIChar) bool {
		return value == "" || value == EmptyChar
	}

	if isEmpty(d.Fill) && isEmpty(d.Outline) {
		return errors.Error("at least one value must be informed to fill or outline")
	}

	if err := d.Fill.Validate(); err != nil {
		return err
	}

	if err := d.Outline.Validate(); err != nil {
		return err
	}

	if d.X < 0 || d.Y < 0 {
		return errors.Error("coordinates must be equal or greater than zero")
	}

	if d.Width <= 0 || d.Height <= 0 {
		return errors.Error("width and height must be equal or greater than zero")
	}

	return nil
}

func (d DrawRequest) IsLateralOutline(column int) bool {
	return column == d.X || column == d.WidthEnd()-1
}
