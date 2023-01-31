//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"unicode"
)

func AnalyzeWord(input string) int {
	count := 0
	for _, l := range input {
		if unicode.IsLetter(l) {
			count += 1
		}
	}
	return count
}

type TotalLetters struct {
	count int
	sync.Mutex
}

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	var wg sync.WaitGroup
	total := 0

	// using a channel and wait groups; no mutexes
	// totalChan := make(chan int, 10)
	// for r.Scan() {
	// 	input := r.Text()
	// 	wg.Add(1)
	// 	go func(totalChan chan int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		totalChan <- AnalyzeWord(input)
	// 	}(totalChan, &wg)
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(totalChan)
	// }()
	// for n := range totalChan {
	// 	total += n
	// }

	// fmt.Println("Total Number of Letters:", total)

	// using mutex and wait groups without a channel
	totalLetters := TotalLetters{}
	for {
		if r.Scan() {
			input := r.Text()
			wg.Add(1)
			go func(word string) {
				totalLetters.Lock()
				defer totalLetters.Unlock()
				defer wg.Done()
				sum := AnalyzeWord(word)
				totalLetters.count += sum
			}(input)
		} else {
			break
		}
	}

	wg.Wait()

	totalLetters.Lock()
	total = totalLetters.count
	totalLetters.Unlock()

	fmt.Println("Total Number of Letters:", total)
}
