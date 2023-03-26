package main

import "fmt"

func do(m1 *map[int]int) {
	(*m1)[3] = 0
	// LOCAL VARIABLE in the function,
	// initialized with a value passed in the actual parameters when the function is called
	*m1 = make(map[int]int) // LOCAL VARIABLE in the function,
	(*m1)[4] = 4
	fmt.Println("m1: ", *m1)
}

func main() {
	m := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("m", m)
	do(&m) // maps are passed by reference
	fmt.Println("m: ", m)
}
