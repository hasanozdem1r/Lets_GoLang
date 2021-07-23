package main

import (
	"fmt"
)

//initial function > main
func main() {
	fmt.Println("Hello Go!")
	fmt.Println(greetings_msg("love yourself"))
	fmt.Println(simple_calculator("+", 12, 1231))
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
