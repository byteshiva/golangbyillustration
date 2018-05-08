package main

import (
	"fmt"
	"math"
)

// Vertex struct definition for X, Y of float64
type Vertex struct {
	X, Y float64
}

// Abs absolute method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale method for scaling X and Y
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
