package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Line struct {
	from  Point
	to    Point
	color rl.Color
	width float32
}

func NewLine(from, to Point, color rl.Color, width float32) *Line {
	return &Line{from: from, to: to, color: color, width: width}
}

func (l *Line) Draw() {
	rl.DrawLineEx(l.from.Vector2, l.to.Vector2, l.width, l.color)
}
