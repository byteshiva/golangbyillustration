package main

import (
	"fmt"
)

//. Example illustrates slice literals
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{4, false},
		{5, true},
		{8, false},
		{5, true},
		{12, false},
	}
	fmt.Println(s)
}
