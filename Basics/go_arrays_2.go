package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s)
	// append works on nil slices.
	s = append(s, 0)
	fmt.Println(s)
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	fmt.Println(s)
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}
