package main

import (
	"fmt"
	"math"
)

// Point represents a point in space
type Point struct {
	X int
	Y int
}

// PointThree represent a three dimentional in space
type PointThree struct {
	Point
	Z int
}

// New returns a Point based on X and Y positions on a graph.
func New(x int, y int) Point {
	return Point{x, y}
}

// NewByThree returns a Point based on X and Y, Z positions on a graph.
func NewByThree(p Point, z int) PointThree {
	return PointThree{p, z}
}

// Distance finds the length of the hypotenuse between two points.
// Formula is the square root of (x2 - x1)^2 + (y2 - y1)^2
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}

// DistanceByThree finds the length of the hypotenuse between two points.
// Formula is the cube root of (x2 - x1)^3 + (y2 - y1)^3 + (z2 - z1)^3
func (p PointThree) DistanceByThree(p2 PointThree) float64 {
	first := math.Pow(float64(p2.X-p.X), 3)
	second := math.Pow(float64(p2.Y-p.Y), 3)
	third := math.Pow(float64(p2.Z-p.Z), 3)
	return math.Cbrt(first + second + third)
}

// main is the entry point for the application.
func main() {
	p1 := New(38, -7)
	p2 := New(20, -80)

	p31 := NewByThree(p1, -12)
	p32 := NewByThree(p2, -92)

	dist := p1.Distance(p2)
	fmt.Println("Distance", dist)

	distby3 := p31.DistanceByThree(p32)
	fmt.Println("Distance", distby3)
}
