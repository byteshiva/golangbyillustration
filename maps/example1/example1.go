package main

import (
	"fmt"
)

// Vertex struct definition with attributes Lat, Long with type float32
type Vertex struct {
	Lat, Long float32
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.12312, -71.12312,
	}
	fmt.Println(m["Bell Labs"])
}
