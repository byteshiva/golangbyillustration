package main

import (
	"fmt"
)

func main() {
	var c chan struct{}
	select {
	case <-c:
	case c <- struct{}{}:
	default:
		fmt.Println("Go here.")
	}
}
