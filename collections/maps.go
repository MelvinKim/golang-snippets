package main

import "fmt"

func main() {
	// incorrect memory initialization
	/*
		m := new(map[string]int)
		m["theAnswer"] = 43
		fmt.Println("map values: ", m)
	*/

	// correct initialization
	m := make(map[string]any)
	m["theAnswer"] = 45
	m["firstName"] = "Melvin"
	fmt.Println(m)
}
