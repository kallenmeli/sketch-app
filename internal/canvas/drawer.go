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
		for i := 0; i < request.HeightEnd(); i++ {
			if i < request.Y {
				continue
			}
			for j := 0; j < request.WidthEnd(); j++ {
				// set padding char if index is lower than the border padding
				if j < request.X {
					draw[i][j] = paddingChar
					continue
				}

				// fill the roof with the outline char
				outline := request.GetOutlineChar()
				if outline != "" && i == request.Y {
					draw[i][j] = outline
					continue
				}

				// fill the left and right borders with the outline char
				if outline != "" &&
					(i >= request.Y) &&
					(j == request.X || j == request.X+request.Width-1) {
					draw[i][j] = outline
					continue
				}

				// fill the footer with the outline
				if outline != "" && i == request.HeightEnd()-1 {
					draw[i][j] = request.GetOutlineChar()
					continue
				}

				if i >= request.Y && i < request.Y+request.Height && j >= request.X && j < request.X+request.Width {
					draw[i][j] = request.GetFillChar()
				}
			}
		}
		draws = append(draws, draw)
	}

	return d.drawToString(width, height, draws), nil
}

func (d drawer) drawToString(width int, height int, draws []Draw) string {
	sb := d.joinDraws(width, height, draws)

	result := strings.Builder{}
	for i, row := range sb {
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

func (d drawer) joinDraws(width int, height int, draws []Draw) [][]string {
	result := make([][]string, height)
	for i := 0; i < height; i++ {
		result[i] = make([]string, width)

		for j := 0; j < width; j++ {
			for _, draw := range draws {
				value := draw[i][j]
				currentValue := result[i][j]
				if strings.Trim(value, " ") == "" && currentValue != "" {
					continue
				}

				result[i][j] = value
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
