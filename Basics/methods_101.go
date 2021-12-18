package main

import (
	"fmt"
	"math"
)

// Vertex2 create receiver
type Vertex2 struct {
	X, Y float64
}

func (v2 Vertex2) Abs() float64 {
	return math.Sqrt(v2.X*v2.X + v2.Y*v2.Y)
}

// Abs written as a regular function with no change in functionality.
func Abs(v2 Vertex2) float64 {
	return math.Sqrt(v2.X*v2.X + v2.Y*v2.Y)
}

func main() {
	v2 := Vertex2{3, 4}
	v3 := Vertex2{5, 12}
	fmt.Println(v2.Abs())
	fmt.Println(Abs(v3))
}
