package main

import "fmt"

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Printf("User name: %s\n", u.name)
}

func main() {
	u := user{name: "Bill"}

	// add the values and pointers to the slice of
	// printer value interfaces
	entities := []printer{
		u,
		&u,
	}

	u.name = "Melvin"
	fmt.Println("User name: ", u.name)

	for _, entry := range entities {
		fmt.Println("entry name: ", entry)
	}
}
