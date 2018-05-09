package main

import (
	"fmt"
)

// Interfaces are implemented implicitly

// I defines Interface
type I interface {
	M()
}

// T struct definition with S of type String
type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so/.

// M prints string in the struct
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
