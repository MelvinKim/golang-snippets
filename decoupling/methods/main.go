package main

import "fmt"

type user struct {
	name  string
	email string
}

// value receiver
func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}

// pointer receiver
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// values of type user can be used  to call methods
	// declared with both value and pointer receivers
	bill := user{"Bill", "bill@gmail.com"}
	bill.changeEmail("bill@hotmail.com")
	bill.notify()

	// pointers of type user can also be used to call methods
	// declared with both value and pointer receivers
	joan := &user{"Joan", "joan@gmail.com"}
	joan.changeEmail("joan@yahoo.com")
	joan.notify()

	users := []user{
		{"ed", "ed@gmail.com"},
		{"erick", "erick@gmail.com"},
	}
	fmt.Println("users: ", users)

	// Not Good! Antipatterns
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}
	fmt.Println("users: ", users)
}

/*
- methods allow data to have behavior
- when should a piece of data have behavior:
- value and pointer semantics in Go
- Other receivers in other languages
1. this (Java)
2. self(Python)
- the receiver name, should never be more that 3 characters eg --> (u user) | (c car) | (r repository)

NB:
- Pointer receivers --> share values
- value receivers --> copy value


*/
