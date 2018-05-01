package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behavior
type notifier interface {
	notify()
}

// use defines a user in the program
type user struct {
	name  string
	email string
}

// notify implements a method notifies user
// of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s> \n",
		u.name,
		u.email)
}

// admin represents an admin user with priviledges.
type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "Mike Malon",
			email: "mike@email.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The Embedded inner type's implementation of the
	// interface is "provided" to the outer type.
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
