//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	r := bufio.NewScanner(os.Stdin)

	commandsNum := 0
	linesNum := 0
	blankLineNum := 0

	for r.Scan() {
		input := r.Text()
		if input == "" {
			blankLineNum += 1
			fmt.Println("I can say nothing too")
			fmt.Println()
		} else if strings.ToLower(input) == "q" {
			break
		} else {
			switch input {
			case "Hi":
				fmt.Println("Hi there")
				commandsNum += 1
			case "Bye":
				fmt.Println("See you next time")
				commandsNum += 1
			case "What's the secret to life?":
				fmt.Println("42")
			default:
				fmt.Println("I don't know what to say")
			}
			linesNum += 1
		}
	}
	fmt.Println("... shutting down ...")
	fmt.Println()
	fmt.Println("Program stats:")
	fmt.Printf("Commands ran: %v | Lines entered: %v | Blank Lines: %v\n", commandsNum, linesNum, blankLineNum)
	fmt.Println()

}
