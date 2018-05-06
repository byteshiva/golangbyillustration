package main

import (
	"fmt"
)

// fibonacci is a funcion that returns
// a funciton that returns an int.
func fibonacci() func() int {
	count := 0
	first, second := 0, 1
	return func() int {
		count++
		fmt.Println("count ", count)
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
