package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	tl      Point
	br      Point
	cntr    Point
	walls   []Side
	wlines  []*Line
	visited bool
	i       int
	next    []*Cell
}

type Side = int8

const (
	WTOP Side = iota
	WBOT
	WLEFT
	WRIGHT
)

type CellPair struct {
	other  *Cell
	s1, s2 Side
}

var W_ALL = []Side{WTOP, WBOT, WLEFT, WRIGHT}

func NewCell(x1, y1, x2, y2 float32, walls []Side) *Cell {
	c := &Cell{tl: *NewPoint(x1, y1), br: *NewPoint(x2, y2), walls: walls}
	c.cntr = *MiddlePoint(c.tl, c.br)
	c.CalcWlines()
	return c
}

func (c *Cell) SetIndex(i int) *Cell {
	c.i = i
	return c
}

func (c *Cell) SetWalls(walls []Side) {
	c.walls = walls
	c.CalcWlines()
}

func (c *Cell) DelWall(wall Side) {
	walls := make([]Side, 0)
	for _, other := range c.walls {
		if other != wall {
			walls = append(walls, other)
		}
	}

	c.walls = walls
	c.CalcWlines()
}

func (c *Cell) CalcWlines() {
	wlines := make([]*Line, 0)

	for _, w := range c.walls {
		switch w {
		case WLEFT:
			wlines = append(wlines, NewLine(
				c.tl, *NewPoint(c.tl.X, c.br.Y),
				rl.Black, 4.0))
		case WRIGHT:
			wlines = append(wlines, NewLine(
				*NewPoint(c.br.X, c.tl.Y), c.br,
				rl.Black, 4.0))
		case WTOP:
			wlines = append(wlines, NewLine(
				c.tl, *NewPoint(c.br.X, c.tl.Y),
				rl.Black, 4.0))
		case WBOT:
			wlines = append(wlines, NewLine(
				*NewPoint(c.tl.X, c.br.Y), c.br,
				rl.Black, 4.0))
		}
	}

	c.wlines = wlines
}

func (c *Cell) Draw() {
	for _, wline := range c.wlines {
		wline.Draw()
	}
}

func (c *Cell) DrawMove(dest Cell, undo bool) {
	NewLine(c.cntr, dest.cntr, rl.Red, 4).Draw()
}

func (c *Cell) BreakPair(cp CellPair) {
	c.DelWall(cp.s1)
	cp.other.DelWall(cp.s2)
}
