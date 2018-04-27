package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email to %s <%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func (u user) emailChangeNotify() {
	fmt.Printf(">>> Email updated to <%s> \n", u.email)
}

func main() {

	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail("bill@gruntemail.com")
	bill.emailChangeNotify()

	joan := &user{"Joan", "joan@email.com"}
	joan.notify()
	joan.changeEmail("joan@gruntemail.com")
	joan.emailChangeNotify()

	// Create a slice of user values with two users
	users := []user{
		{"bill", "bill@email.com"},
		{"joan", "joan@email.com"},
	}

	// Iterate over the slice of users switching
	for _, u := range users {
		u.changeEmail("it@wontmatternow.com")
		u.emailChangeNotify()
	}
}
