package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	cx := p.x - other.x
	cy := p.y - other.y
	return math.Sqrt(cx*cx + cy*cy)
}

func main() {

	p1 := NewPoint(-4.0, 2.0)
	p2 := NewPoint(7.0, -6.0)

	c := p1.Distance(p2)

	fmt.Printf("Расстояние между точками: %.2f\n", c)
}
