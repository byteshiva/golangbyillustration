package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name     string
	age      int
	building float32
	secure   bool
	roles    []string
}

func main() {
	// Create a struct type user value.
	u := user{
		name:     "Cindy",
		age:      27,
		building: 321.44,
		secure:   true,
		roles:    []string{"admin", "developer", "customer"},
	}

	// Store a copy of the user value inside an empty interface value.
	var i interface{} = u

	// Display information about the user value that was stored.
	v := reflect.ValueOf(i)
	fmt.Printf("Kind: %v\t Type: %v\t\t NumFields: %v\n", v.Kind(), v.Type(), v.NumField())

}
