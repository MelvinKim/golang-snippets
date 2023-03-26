package main

import "fmt"

func doArray(b [3]int) int {
	// changes made here are  ONLY LOCAL TO THIS FUNCTION
	// Arrays are passed by value
	b[0] = 0
	return b[1]
}

func doSlice(b []int) int {
	// changes made here are GLOBAL
	// Slices are passed by reference --> hence they will persist
	b[0] = 0
	fmt.Printf("b@ %p\n", b) // prints the meory address
	return b[1]
}

func main() {
	aArray := [3]int{1, 2, 3}
	aSlice := []int{1, 2, 3}
	fmt.Printf("a@ %p\n", aSlice) // prints the memory address
	vArray := doArray(aArray)
	vSlice := doSlice(aSlice)

	fmt.Println(aArray, vArray)
	fmt.Println(aSlice, vSlice)
}
