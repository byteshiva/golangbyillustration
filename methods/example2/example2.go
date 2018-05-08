package main

import (
	"fmt"
	"math"
)

// Vertex  Struct definition for X ,Y of type float
type Vertex struct {
	X, Y float64
}

// Abs absolute function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
