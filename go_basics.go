package main

import (
	"fmt"
	"math/cmplx"
)

//initial function > main
func main() {
	var (
		//variable declaration
		name, surname, nationality string
		//variable declaration with initializers
		birthday, personal_id int = 1997, 1234567890
		//basic data types
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	var (
		x_int   = 12
		y_float = 12.1
	)
	const ID = 1212

	//short variable declaration with variable_name:=variable_value
	decision := true
	//type conversion ways
	fmt.Println(x_int, y_float)
	x_int = int(y_float)
	y_float = float64(x_int)
	fmt.Println(x_int, y_float)
	fmt.Println(name, surname, nationality)
	fmt.Println(birthday, personal_id)
	fmt.Println(decision)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("Hello Go!")
	//call greetings_msg function
	fmt.Println(greetings_msg("love yourself"))
	//call simple_calculator function
	fmt.Println("Result:", simple_calculator("+", 12, 1231))
}

func greetings_msg(msg string) string {
	msg_edited := fmt.Sprintf("Hi, here is my message to you > %v <", msg)
	return msg_edited
}

func simple_calculator(operand string, number1, number2 float32) float32 {
	if operand == "+" {
		return number1 + number2
	} else if operand == "-" {
		return number1 - number2
	} else if operand == "*" {
		return number1 * number2
	} else {
		return number1 / number2
	}
}
