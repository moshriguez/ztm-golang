//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type UnicodeFn func(r rune) bool

var runeTypeToFn = map[string]UnicodeFn{
	"letter": unicode.IsLetter,
	"digit": unicode.IsDigit,
	"space": unicode.IsSpace,
	"punctuation mark": unicode.IsPunct, 
}

func ProcessLineOfText(lines []string, runeType string, processFn func(input []string, runeType string)) {
	processFn(lines, runeType)
}

func PrintNumOf(input []string, runeType string) {
	count := 0
	fn, ok := runeTypeToFn[runeType]
	if !ok {
		panic("We don't know what function to run for that runeType")
	}
	for _, line := range input {
		for _, rune := range line {
			if fn(rune) {
				count += 1
			}
		}
	}

	fmt.Printf("The text has %v %v%v\n", count, runeType, plural(count))
}

func plural(count int) string {
	if count == 1 {
		return ""
	} else {
		return "s"
	}
}

// func PrintNumOfLetters(input []string) {
// 	numOfLetters := 0
// 	for _, line := range input {
// 		for _, rune := range line {
// 			if unicode.IsLetter(rune) {
// 				numOfLetters += 1
// 			}
// 		}
// 	}
// 	fmt.Printf("The text has %v letters\n", numOfLetters)
// }

// func PrintNumOfDigits(input []string) {
// 	numOfDigits := 0
// 	for _, line := range input {
// 		for _, rune := range line {
// 			if unicode.IsDigit(rune) {
// 				numOfDigits += 1
// 			}
// 		}
// 	}
// 	fmt.Printf("The text has %v digits\n", numOfDigits)
// }

// func PrintNumOfSpaces(input []string) {
// 	numOfSpaces := 0
// 	for _, line := range input {
// 		for _, rune := range line {
// 			if unicode.IsSpace(rune) {
// 				numOfSpaces += 1
// 			}
// 		}
// 	}
// 	fmt.Printf("The text has %v spaces\n", numOfSpaces)
// }

// func PrintNumOfPuncuation(input []string) {
// 	numOfPunct := 0
// 	for _, line := range input {
// 		for _, rune := range line {
// 			if unicode.IsPunct(rune) {
// 				numOfPunct += 1
// 			}
// 		}
// 	}
// 	fmt.Printf("The text has %v puncuation marks\n", numOfPunct)

// }

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}
	ProcessLineOfText(lines, "letter", PrintNumOf)
	ProcessLineOfText(lines, "digit", PrintNumOf)
	ProcessLineOfText(lines, "space", PrintNumOf)
	ProcessLineOfText(lines, "punctuation mark", PrintNumOf)
}
