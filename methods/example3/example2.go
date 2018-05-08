package main

import (
	"fmt"
	"math"
)

// MyFloat of float 64 bit type
type MyFloat float64

// Abs absolute definition
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
