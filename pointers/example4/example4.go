package main

import (
	"fmt"
	"math"
)

// Vertex struct definition with X, Y of float 64 bit
type Vertex struct {
	X, Y float64
}

// Abs Method returns math square root of X, Y
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AbsFunc Method returns math square root of X, Y
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
