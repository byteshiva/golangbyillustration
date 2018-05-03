package main

import (
	"fmt"
)

// Vertex - defines fields with X , Y with type integer, Collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{23, 123})
}
