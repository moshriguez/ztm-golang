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

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Split(bufio.ScanWords)

	var wg sync.WaitGroup
	totalChan := make(chan int, 10)
	total := 0
	for r.Scan() {
		input := r.Text()
		wg.Add(1)
		go func (totalChan chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			totalChan <- AnalyzeWord(input)
		}(totalChan, &wg)
		
	}
	go func() {
		wg.Wait()
		close(totalChan)
	}()
	for n := range totalChan {
		total += n
	}
	fmt.Println("Total Number of Letters:", total)
}
