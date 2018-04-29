package main

import (
	"fmt"
)

// Animal contains all the base fields for animals.
type Animal struct {
	Name string
}

// Speaker provide a common behavior for all concrete types
// to follow if they want to be a part of this group. This is
// a contract for these concrete types to follow.
type Speaker interface {
	Speak()
}

// // Speak provides generic behavior for all animals and how they speak
// func (a *Animal) Speak() {
// 	fmt.Println("UGH!",
// 		"My name is", a.Name,
// 		", it is", a.IsMammal,
// 		"I am a mammal")
// }

//Dog contains everything an Animal is but specific
// attributes that only a Dog has.
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak knows how to speak like a dog.
func (d *Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat contains everything an Animal is but specific
// attributes that only a Cat has.
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// Speak kows how to speak like a Cat
func (c *Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is ", c.IsMammal,
		"I am a mammal with a climb facotr of", c.ClimbFactor)
}

func main() {
	// Create a list of Animals that how how to Speak
	// and then its specific Dog attributes.
	speakers := []Speaker{
		&Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},
		// Create a Cat by initializing its Animal parts
		// and then its specific Cat attributes.
		&Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	// Have the Animals Speak.
	for _, spkr := range speakers {
		spkr.Speak()
	}
}
