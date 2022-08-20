package canvas

import "time"

type Canvas struct {
	ID        string    `json:"id"`
	Drawing   string    `json:"drawing"`
	CreatedAt time.Time `json:"created_at"`
}

type DrawRequest struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Outline string `json:"outline"`
	Fill    string `json:"fill"`
}

func (d DrawRequest) GetFillChar() string {
	if d.Fill == "none" {
		return " "
	}
	return d.Fill
}

func (d DrawRequest) GetOutlineChar() string {
	if d.Outline == "none" {
		return ""
	}
	return d.Outline
}

func (d DrawRequest) WidthEnd() int {
	return d.X + d.Width
}

func (d DrawRequest) HeightEnd() int {
	return d.Y + d.Height
}

type DrawResponse struct {
	ID      string `json:"id"`
	Drawing string `json:"canvas"`
}
