package main

import "fmt"

// user defines a user in the system
type user struct {
	name  string
	email string
}

// user defines a user in the system
type member struct {
	name  string
	email string
}

// String implements the fmt.Stringer interface.
func (u *user) String() string {
	return fmt.Sprintf("My Name is %q and my email is %q", u.name, u.email)
}

// String implements the fmt.Stringer interface.
func (u member) String() string {
	return fmt.Sprintf("My Name is %q and my email is %q", u.name, u.email)
}

func main() {
	// create a value of type user.

	u := user{
		name:  "Bill",
		email: "bill@testemailonly.com",
	}

	m := member{
		name:  "John",
		email: "John@testmail.com",
	}

	// Display the values.
	fmt.Println(u)
	fmt.Println(&u)

	// Display the values.
	fmt.Println(m)
	fmt.Println(&m)

	fmt.Println()

}
