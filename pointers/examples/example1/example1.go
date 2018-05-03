package main

import (
	"fmt"
)

func main() {
	i, j, k := 42, 1278, 34

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // set the new value i

	p = &j         // point to j
	*p = *p / 34   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	p = &k
	*p = *p * 89
	fmt.Println(k)
}
