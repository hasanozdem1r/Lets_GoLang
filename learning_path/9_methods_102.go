package main

import (
	"fmt"
	"math"
)

// Vertex3 is a name of struct
type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
