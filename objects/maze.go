package objects

import (
	"math/rand"
	"slices"
)

type Maze struct {
	tl         *Point
	rows, cols int
	cellAmount int
	cellSize   float32
	cells      []*Cell
	rand       *rand.Rand
}

func NewMaze(x, y float32, rows, cols int, cellSize float32, seed int64) *Maze {
	m := &Maze{
		tl:         NewPoint(x, y),
		rows:       rows,
		cols:       cols,
		cellAmount: rows * cols,
		cellSize:   cellSize,

		rand: rand.New(rand.NewSource(seed)),
	}
	m.createCells()

	return m
}

func (m *Maze) createCells() {
	var i, j float32
	var index int
	for i = 0; i < float32(m.rows); i++ {
		for j = 0; j < float32(m.cols); j++ {
			m.cells = append(m.cells, NewCell(
				m.tl.X+j*m.cellSize, m.tl.Y+i*m.cellSize,
				m.tl.X+(j+1)*m.cellSize, m.tl.Y+(i+1)*m.cellSize,
				W_ALL).SetIndex(index))
			index++
		}
	}
	m.breakEntranceAndExit()
	m.breakWalls(m.cells[0])
	m.resetVisits()
}

func (m *Maze) breakEntranceAndExit() {
	m.cells[0].SetWalls([]Side{WTOP, WBOT, WRIGHT})
	m.cells[m.cellAmount-1].SetWalls([]Side{WTOP, WBOT, WLEFT})
}

func (m *Maze) breakWalls(c *Cell) {
	c.visited = true
	cpairs := m.findNeighbours(c)
	for len(cpairs) > 0 {
		i := m.rand.Intn(len(cpairs))
		cpair := cpairs[i]

		cpairs = slices.Delete(cpairs, i, i+1)
		if cpair.other.visited {
			continue
		}
		c.BreakPair(cpair)
		c.next = append(c.next, cpair.other)
		m.breakWalls(cpair.other)
	}
}

func (m *Maze) resetVisits() {
	for _, cell := range m.cells {
		cell.visited = false
	}
}

func (m *Maze) findNeighbours(c *Cell) []CellPair {
	cnext := make([]CellPair, 0)
	if c.i%m.rows != 0 && !m.cells[c.i-1].visited {
		cnext = append(cnext, CellPair{
			other: m.cells[c.i-1],
			s1:    WLEFT, s2: WRIGHT,
		})
	}
	if c.i%m.rows != m.rows-1 && !m.cells[c.i+1].visited {
		cnext = append(cnext, CellPair{
			other: m.cells[c.i+1],
			s1:    WRIGHT, s2: WLEFT,
		})
	}
	if c.i >= m.cols && !m.cells[c.i-m.cols].visited {
		cnext = append(cnext, CellPair{
			other: m.cells[c.i-m.cols],
			s1:    WTOP, s2: WBOT,
		})
	}
	if c.i < m.rows*(m.cols-1) && !m.cells[c.i+m.cols].visited {
		cnext = append(cnext, CellPair{
			other: m.cells[c.i+m.cols],
			s1:    WBOT, s2: WTOP,
		})
	}
	return cnext
}

func (m *Maze) Draw() {
	for _, cell := range m.cells {
		cell.Draw()
	}
}

func (m *Maze) Solve() bool {
	return m.solveRec(m.cells[0])
}

func (m *Maze) solveRec(c *Cell) bool {

	return false
}
