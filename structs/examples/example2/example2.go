package main

import (
	"fmt"
)

// Vertex Struct - defines X, Y of type integer.
type Vertex struct {
	X int
	Y int
}

func main() {
	// Struct Initialization
	v := Vertex{1, 2}
	// Vertex struct are accessed using a dot notation.
	v.X = 4
	fmt.Println(v.X)
}
