package model

import "math"

type Position [2]int

func NewPosition(x, y int) *Position {
	return &Position{x, y}
}

func (p *Position) Inside(pos *Position, radius float64) bool {
	x := float64(p[0] - pos[0])
	x = math.Pow(x, 2.0)

	y := float64(p[1] - pos[1])
	y = math.Pow(y, 2.0)

	return radius > math.Sqrt(x+y)
}
