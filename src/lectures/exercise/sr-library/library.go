//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	title string
	author string
	totalCopies int
	numCheckedOut int
}

type CheckedOutBook struct {
	checkedOutDate time.Time
	returnedDate time.Time
}

type Member struct {
	name string
	borrowed map[string]CheckedOutBook
}

type Library struct {
	Books map[string]Book
	Members map[string]Member
}

func AddBook(lib *Library, title, author string, number int) {
	book := Book{
		title: title,
		author: author,
		totalCopies: number,
	}
	lib.Books[title] = book
}

func AddMember(lib *Library, name string) {
	member := Member{name, make(map[string]CheckedOutBook)}
	lib.Members[name] = member
}

func PrintLibrary(lib *Library) {
	fmt.Println("- - - - - - -")
	fmt.Println("Starting Library")
	fmt.Println("- - - - - - -")
	fmt.Println("- Members -")
	fmt.Println("- - - - - - -")
	for _, member := range lib.Members {
		fmt.Println(member.name)
		fmt.Println()
	}
	fmt.Println("- - - - - - -")
	fmt.Println("- Books -")
	for _, book := range lib.Books {
		fmt.Println(book.title, "by", book.author)
		fmt.Println("total copies:", book.totalCopies, "/ copies available:", book.totalCopies - book.numCheckedOut)
		fmt.Println()
	}
}

func CheckOutBook(lib *Library, title, memberName string) error {
	book, ok := lib.Books[title]
	if !ok {
		return fmt.Errorf("book not found in library")
	}
	if book.numCheckedOut == book.totalCopies {
		return fmt.Errorf("all copies are checked out")
	}
	member, ok := lib.Members[memberName]
	if !ok {
		return fmt.Errorf("member not found in library")
	}
	member.borrowed[title] = CheckedOutBook{
		checkedOutDate: time.Now(),
	}
	book.numCheckedOut += 1
	lib.Books[title] = book
	lib.Members[memberName] = member

	return nil
}

func CheckInBook(lib *Library, title, memberName string) error {
	book, ok := lib.Books[title]
	if !ok {
		return fmt.Errorf("book not found in library")
	}
	member, ok := lib.Members[memberName]
	if !ok {
		return fmt.Errorf("member not found in library")
	}
	checkedOutBook, ok := member.borrowed[title]
	if !ok {
		return fmt.Errorf("book was not checked out by this member")
	}
	book.numCheckedOut -= 1
	checkedOutBook.returnedDate = time.Now()
	member.borrowed[title] = checkedOutBook
	
	lib.Books[title] = book
	lib.Members[memberName] = member

	return nil
}

func main() {
	lib := Library{make(map[string]Book), make(map[string]Member)}
	AddBook(&lib, "One Fish, Two Fish, Red Fish, Blue Fish", "Dr. Suess", 2)
	AddBook(&lib, "Pachinko", "Min Jin Lee", 5)
	AddBook(&lib, "It", "Stephen King", 3)
	AddBook(&lib, "Post Office", "Charles Bukowski", 1)

	AddMember(&lib, "Charles Darwin")
	AddMember(&lib, "Michael Jordan")
	AddMember(&lib, "Bob Dylan")

	PrintLibrary(&lib)
	err := CheckOutBook(&lib, "One Fish, Two Fish, Red Fish, Blue Fish", "Bob Dylan")
	if err != nil {
		fmt.Println(err)
	}
	PrintLibrary(&lib)
	err = CheckOutBook(&lib, "It", "Michael Jordan")
	if err != nil {
		fmt.Println(err)
	}
	PrintLibrary(&lib)
	err = CheckInBook(&lib, "It", "Michael Jordan")
	if err != nil {
		fmt.Println(err)
	}
	PrintLibrary(&lib)

}
