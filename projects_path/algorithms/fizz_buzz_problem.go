package main

import "fmt"

func main() {
	fizz_buzz(50)
}

// this method is created as a solution of fizz_buzz problem
func fizz_buzz(number int) {
	// loop definition
	for i := 0; i < number; i++ {
		// fizbuzz condition
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
			// fizz condition
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
			// buzz condition
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
			// no matching pattern
		} else {
			fmt.Println("None")
		}
	}
}
