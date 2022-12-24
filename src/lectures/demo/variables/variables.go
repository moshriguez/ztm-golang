package main

import "fmt"

func main() {
	//variable declaration and assignment
	var myName = "Marc"

	fmt.Println("my name is", myName)

	// variable declaration and assignment with type annotation
	var name string = "Kathy"
	fmt.Println("name =", name)

	// shorthand declaration/assignment (create and assign) (uses type inference)
	userName := "admin"
	fmt.Println("username =", userName)

	// uninitialized variable - declaration without assignment (variable has default value)
	var sum int
	fmt.Println("The sum is", sum)

	// compound assignment
	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	// ok comma idiom - you can reuse the 2nd variable
	// commonly used for errors or checking if a fields is present in a map
	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	// reassign a variable
	sum = part1 + part2
	fmt.Println("sum =", sum)

	// block assignment
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	// use _ to ignore a variable
	// below is not a very common use
	// more commonly, you would use this if a function has multiple return values, but you dont need all return values
	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
