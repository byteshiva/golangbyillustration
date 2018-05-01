package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

// notify implements a methods notifies users of different events.
func (u *user) notify(level string) {
	if level != "" {
		fmt.Printf("Sending user email.. to %s<%s> ..\n",
			u.name,
			u.email)
	} else {
		fmt.Printf("Sending user email to %s<%s>\n",
			u.name,
			u.email)
	}
}

// admin represents an admin user with privileges.
type admin struct {
	person user
	level  string
}

func main() {
	// Create an admin user.
	ad := admin{
		person: user{
			name:  "Mike menkon",
			email: "mike@email.com",
		},
		level: "super",
	}

	// We can access fields methods.
	ad.person.notify("test")
}
