package main

import "fmt"

type I interface {
	M()
}

// Notice, i have not specified a type to implement the method in the interface, so (VALUE, TYPE == nil)
func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
