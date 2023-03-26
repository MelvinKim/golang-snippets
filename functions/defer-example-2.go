package main

import "fmt"

// defer statement runs before the return is "done"
func doIt() (a int) {
	defer func() {
		a = 2
		fmt.Println("a inside defer: ", a)
	}()

	a = 1
	fmt.Println("a outside defer: ", a)
	return
}

func main() {

	doIt()

	a := 10
	defer fmt.Println(a) // 10 : defer copies arguments to the deferred call

	a = 11

	fmt.Println(a) // 11
}
