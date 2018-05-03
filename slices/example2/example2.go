package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"John",
		"Pael",
		"Gering",
		"Ngino",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXXYYY"
	fmt.Println(a, b)
	fmt.Println(names)
}
