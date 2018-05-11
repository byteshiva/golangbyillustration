package main

import (
	"fmt"
)

// Person defines two types Name and Age
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
	a := Person{" Arthor", 33}
	z := Person{"Zapod Beebnlebox", 9901}
	y := Person{"Bamboo Box", 91}
	fmt.Println(a, z, y)
}
