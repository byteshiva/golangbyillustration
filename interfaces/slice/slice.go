package main

import (
	"fmt"
	"reflect"
)

// user represents a basic user in the system.
type user struct {
	name     string
	age      int
	building float32
	secure   bool
	roles    []string
}

func main() {
	us := []user{
		{
			name:     "Cindy",
			age:      27,
			building: 33.44,
			roles:    []string{"admin", "developer"},
		},
		{
			name:     "Bill",
			age:      40,
			building: 321.22,
			secure:   false,
			roles:    []string{"developer"},
		},
		{
			name:   "Mike",
			age:    35,
			secure: true,
			roles:  []string{"admin"},
		},
	}

	// Store a value of the slice inside an empty interface value.
	var i interface{} = us

	// Display information about the slice stored inside the empty
	// interface value.
	v := reflect.ValueOf(i)
	fmt.Printf("Kind: %v\t Type: %v\n", v.Kind(), v.Type())

	// Iterate over the slice via reflection.
	for i := 0; i < v.Len(); i++ {
		fmt.Println(v.Index(i))
	}

}
