package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	fmt.Println("Array values: ", colors)

	numbers := [3]int{1, 2, 4}
	fmt.Println("numbers array: ", numbers)
	fmt.Println("number array length: ", len(numbers))
}
