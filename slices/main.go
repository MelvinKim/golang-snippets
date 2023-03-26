package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Println("length (values): ", len(values))
	fmt.Println("capacity (values): ", cap(values))
	fmt.Println("---------------------------------")
	slice := values[1:]
	fmt.Println("length (slice): ", len(slice))
	fmt.Println("capacity (slice): ", cap(slice))
	fmt.Println("length (values): ", len(values))
	fmt.Println("capacity (values): ", cap(values))
	fmt.Println("---------------------------------")
	slice[0] = 10
	fmt.Println("length (slice): ", len(slice))
	fmt.Println("capacity (slice): ", cap(slice))
	fmt.Println("length (values): ", len(values))
	fmt.Println("capacity (values): ", cap(values))
	fmt.Println("---------------------------------")
	slice = append(slice, 6)
	fmt.Println("length (slice): ", len(slice))
	fmt.Println("capacity (slice): ", cap(slice))
	fmt.Println("length (values): ", len(values))
	fmt.Println("capacity (values): ", cap(values))
	fmt.Println("---------------------------------")
	slice[1] = 20
	fmt.Println("length (slice): ", len(slice))
	fmt.Println("capacity (slice): ", cap(slice))
	fmt.Println("length (values): ", len(values))
	fmt.Println("capacity (values): ", cap(values))

	fmt.Println("---------------------------------")
	fmt.Println("---------------------------------")
	fmt.Printf("%v, len=%d, cap=%d\n", values, len(values), cap(values))
}
