package main

import (
	"fmt"
)

// user defines a user in the program
type user struct {
	name  string
	email string
}

// notify implements a method notifies user of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user  // Embedded type
	level string
}

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "Mike Mayor",
			email: "mike@email.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type is promoted.
	ad.notify()
}
