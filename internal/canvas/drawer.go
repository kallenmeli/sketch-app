package canvas

import (
	"strings"

	"github.com/labstack/gommon/log"
)

const (
	paddingChar = " "
)

type (
	Drawer interface {
		Draw(requests []DrawRequest) (string, error)
	}
	drawer struct {
	}
)

func NewDrawer() Drawer {
	return &drawer{}
}

func (d drawer) Draw(requests []DrawRequest) (string, error) {
	width, height := d.getCanvasDimension(requests)
	log.Infof("width: %v, height: %v", width, height)
	draws := make([]Draw, 0, len(requests))

	if len(requests) == 0 {
		return "", ErrEmptyRequests
	}

	for _, request := range requests {
		draw := NewDraw(width, height)
		for row := request.Y; row < request.HeightEnd(); row++ {
			for column := 0; column < request.WidthEnd(); column++ {

				if column < request.X {
					draw[row][column] = paddingChar
					continue
				}

				if canFill, outline := d.canFillOutline(row, column, request); canFill {
					draw[row][column] = outline
					continue
				}

				draw[row][column] = request.GetFillChar()
			}
		}

		draws = append(draws, draw)
	}

	return d.drawToString(width, height, draws), nil
}

func (d drawer) canFillOutline(row, column int, request DrawRequest) (bool, string) {
	outline := request.GetOutlineChar()

	if outline == "" {
		return false, ""
	}

	if request.IsFirstRow(row) {
		return true, outline
	}

	if row >= request.Y && request.IsLateralOutline(column) {
		return true, outline
	}

	if request.IsLastRow(row) {
		return true, outline
	}

	return false, ""
}

func (d drawer) drawToString(width int, height int, draws []Draw) string {
	finalDraw := d.joinDraws(width, height, draws)
	return finalDraw.String()
}

func (d drawer) joinDraws(width int, height int, draws []Draw) Draw {
	result := make([][]string, height)
	for row := 0; row < height; row++ {
		result[row] = make([]string, width)

		for column := 0; column < width; column++ {
			for _, draw := range draws {
				value := draw[row][column]
				currentValue := result[row][column]
				cannotBeReplacedWithEmpty := strings.Trim(value, " ") == "" && currentValue != ""
				if cannotBeReplacedWithEmpty {
					continue
				}

				result[row][column] = value
			}
		}
	}
	return result
}

func (d drawer) getCanvasDimension(requests []DrawRequest) (int, int) {
	width := 0
	height := 0

	for _, request := range requests {
		currentWidth := request.X + request.Width
		if currentWidth > width {
			width = currentWidth
		}
		if request.HeightEnd() > height {
			height = request.HeightEnd()
		}
	}

	return width, height
}
