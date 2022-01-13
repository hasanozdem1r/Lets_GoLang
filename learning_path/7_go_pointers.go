package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i // point to i
	fmt.Println(*p)

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)

	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{12, 23}
	fmt.Println(Vertex{1, 2})
	fmt.Println(v.X, v.Y)
	// struct literals
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		k  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, k, v2, v3)

	// array declaration
	var name [3]string
	name[0] = "Hasan"
	name[1] = "Anastasia"
	name[2] = "Our_Baby"

	singular := [4]int{1, 3, 5, 7}
	for item := range singular {
		fmt.Println(item)
	}

}
