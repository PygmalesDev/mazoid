package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Point struct {
	rl.Vector2
}

func NewPoint(x, y float32) *Point {
	return &Point{rl.Vector2{X: x, Y: y}}
}

func MiddlePoint(p1, p2 Point) *Point {
	return NewPoint(p1.X+(p2.X-p1.X)/2, p1.Y+(p2.Y-p1.Y)/2)
}
