package main

import "fmt"

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("pointer value of P: ", p)
	fmt.Println("Integer value of p: ", *p)

	aFloat := 42.13
	var f = &aFloat
	fmt.Println("pointer value of f: ", f)
	fmt.Println("Integer value of f: ", *f)

	*f = *f / 31
	fmt.Println("Pointer: ", *f)
	fmt.Println("Value 1: ", aFloat)
	// pointer variable,
	// if not assigned a value, the intial value will be nil
	/*
		var p *int
		fmt.Println(*p)
	*/
}
