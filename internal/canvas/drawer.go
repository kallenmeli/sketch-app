package canvas

import (
	"github.com/labstack/gommon/log"
)

const (
	borderPadding = 2
	paddingChar   = " "
)

type (
	Drawer struct {
	}
	Draw [][]string
)

func (d Drawer) Draw(requests []DrawRequest) string {
	width, height := d.getCanvasDimension(requests)
	log.Infof("width: %v, height: %v", width, height)
	draws := make([]Draw, 0, len(requests))

	for _, request := range requests {
		draw := d.createEmptyDraw(height, width)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				// set padding char if index is lower than the border padding
				if j < request.X {
					draw[i][j] = paddingChar
					continue
				}

				// fill the roof with the outline char
				if request.Outline != "" && i == request.Y {
					draw[i][j] = request.Outline
					continue
				}

				// fill the left and right borders with the outline char
				if request.Outline != "" &&
					(i >= request.Y) &&
					(j == request.X || j == request.X+request.Width-1) {
					draw[i][j] = request.Outline
					continue
				}

				// fill the footer with the outline
				if request.Outline != "" && i == request.HeightEnd()-1 {
					draw[i][j] = request.Outline
					continue
				}

				if i >= request.Y && i < request.Y+request.Height && j >= request.X && j < request.X+request.Width {
					draw[i][j] = request.GetFillChar()
				}
			}
		}
		draws = append(draws, draw)
	}

	return d.drawToString(width, height, draws)
}

// transform a draw to a string
func (d Drawer) drawToString(width int, height int, draws []Draw) string {
	sb := make([][]string, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for _, draw := range draws {
				sb[i][j] = draw[i][j]
				//sb.WriteString(draw[i][j])
			}
		}
		isFinalRow := i == height-1
		if !isFinalRow {
			//sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (d Drawer) getCanvasDimension(requests []DrawRequest) (int, int) {
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

func (d Drawer) createEmptyDraw(height, width int) Draw {
	draw := make(Draw, height)
	for i := 0; i < height; i++ {
		draw[i] = make([]string, width)
	}
	return draw
}
