package main

import (
	"fmt"
	"math"
)

// Vertex struct definition with X, Y of float 64 bit
type Vertex struct {
	X, Y float64
}

// Abs returns Absolute value
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale return Scale of factor
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{4, 5}
	Scale(&v, 20)
	fmt.Println(Abs(v))
}
