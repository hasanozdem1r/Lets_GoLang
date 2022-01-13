package main

import "fmt"

func main() {
	// code convention tips
	// camelCase preferred for GoLang
	var name [5]string
	name[0] = "Adam"
	name[1] = "Harry"
	name[2] = "Dan"
	name[3] = "Alex"
	name[4] = "Jack"
	// show all name
	fmt.Println(name)
	// create integer array
	oddNumbers := [5]int{1, 3, 5, 7, 9}
	// show all odd number
	fmt.Println(oddNumbers)

	//SLICING PRACTICES
	// show all
	fmt.Println(oddNumbers[:])
	fmt.Println(oddNumbers[0])
	// structure arrayName[low:high]
	// low index is included, high index is excluded
	fmt.Println(name[2:3])
	// show nothing because 3 is included and 3 is excluded and result None
	fmt.Println(name[3:3])
	// assign value to array by index
	name[1] = "Hasan"
	fmt.Println(name)

	// slice literals
	// A slice literal is like an array literal without the length.
	q := []int{2, 3, 5, 7, 11, 13}
	var q1 []int
	r := []bool{true, false, true, true, false, true}
	fmt.Println(q)
	fmt.Println(r)
	// cap(arr) and len(arr) function return of array
	fmt.Println(cap(q), len(q))

	if len(q1) == 01 {
		fmt.Println("Array is not null")
	} else if q1 == nil {
		fmt.Println("Array is null")
		fmt.Println(len(q1))
	}
	a := make([]int, 5) // len(a)=5
	fmt.Println(a)
	b := make([]int, 1, 5) // len(b)=0, cap(b)=5
	fmt.Println(b)
}
