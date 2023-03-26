package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// (v Vertex) --> receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can also declare a method on non-struct types too
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
