//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Welcome to Golang,", name)
}

//* Write a function that returns any message, and call it from within fmt.Println()
func message() string {
	return "How are you?"
}

//* Write a function to add 3 numbers together, supplied as arguments, and return the answer
func add(x, y, z int) int {
	return x + y + z
}

//* Write a function that returns any number
func number() int {
	return 5
}

//* Write a function that returns any two numbers
func numbers() (int, int) {
	return 3, 7
}

// MAIN Function
func main() {
	greet("Manish Thakur")

	msg := message()
	fmt.Println(msg)

	x := number()

	y, z := numbers()

	sum := add(x, y, z)
	fmt.Println(sum)
}
