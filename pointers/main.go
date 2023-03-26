package main

import "fmt"

func main() {
	x := 4
	xMemoryAddress := &x
	y := &x
	fmt.Println("X memory address: ", &x)
	fmt.Println("Value of Y which stores the memory address of X: ", y)
	fmt.Println("Numeric Value of X, stored in Y: ", *y)

	// copy value of X to a new variable
	xCopy := *y
	xCopyTwo := *xMemoryAddress
	fmt.Println("X copy: ", xCopy)
	fmt.Println("X Copy: ", xCopyTwo)
}
