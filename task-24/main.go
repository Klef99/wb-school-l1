package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(p.X()-other.X(), 2) + math.Pow(p.Y()-other.Y(), 2))
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(5, 6)

	fmt.Printf("Point 1: %v\nPoint 2: %v\n", p1, p2)
	fmt.Println("Distance: ", p1.Distance(*p2))
}
