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
	// Create a map of struct type user values.
	um := map[string]user{
		"Cindy": {
			name:     "Cindy",
			age:      23,
			building: 555.12,
			secure:   true,
			roles:    []string{"admin", "developer"},
		},
		"Bill": {
			name:     "Bill",
			age:      33,
			building: 223.87,
			secure:   false,
			roles:    []string{"developer"},
		},
	}

	// Store a value of the map inside an empty interface value.
	var i interface{} = um

	// Display the information about the map of users values.
	v := reflect.ValueOf(i)
	fmt.Printf("Kind: %v\t Type: %v\n", v.Kind(), v.Type())

	// Iterate over the map via reflection.
	for i, key := range v.MapKeys() {
		fmt.Println(i, ":", v.MapIndex(key))
	}

}
