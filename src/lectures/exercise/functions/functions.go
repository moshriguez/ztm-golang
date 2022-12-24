//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
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

func greetPerson(name string) {
	fmt.Println("Hello", name)
}

func message() string {
	return "this is a sample message"
}

func add3(a, b, c int) int {
	return a + b + c
}

func anyNumber() int {
	return 6
}

func any2Numbers() (int, int) {
	return anyNumber(), 9
}

func main() {
	mike := "Mike"
	greetPerson(mike)

	fmt.Println(message())

	sum := add3(1, 2, 3)
	fmt.Println("3! is", sum)

	fmt.Println(any2Numbers())

	a, b := any2Numbers()
	newSum := add3(a, b, 4)
	fmt.Println("sum of 3 numbers:", newSum)
}
