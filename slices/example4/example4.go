package main

import (
	"fmt"
)

// Slice defaults
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	// s := []int{2}

	// gets the reference values from array[1 to 4]
	s = s[1:4]
	fmt.Println(s)

	s = s[1:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}
