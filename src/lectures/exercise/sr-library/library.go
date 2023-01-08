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
	checkedOut bool
	checkedOutDate time.Time
	returnedDate time.Time
	checkedOutBy	*Member
}

type Member struct {
	name string
}

type Library struct {
	Books []*Book
	Members []*Member
}

func AddBook(lib *Library, title, author string) {
	book := Book{
		title: title,
		author: author,
	}
	lib.Books = append(lib.Books, &book)
}

func AddMember(lib *Library, name string) {
	member := Member{name}
	lib.Members = append(lib.Members, &member)
}

func PrintLibrary(lib *Library) {
	fmt.Println("- - - - - - -")
	fmt.Println("Starting Library")
	fmt.Println("- - - - - - -")
	fmt.Println("- Members -")
	fmt.Println("- - - - - - -")
	for _, member := range lib.Members {
		fmt.Println(member.name)
		fmt.Println("")
	}
	fmt.Println("- - - - - - -")
	fmt.Println("- Books -")
	for _, book := range lib.Books {
		fmt.Println(book.title, "by", book.author)
		fmt.Println("checked out:", book.checkedOut)
		if book.checkedOut {
			fmt.Println("checked out by:", book.checkedOutBy.name)
		}
		fmt.Println("")
	}
}

func CheckOutBook(lib *Library, title, memberName string) error {
	book, err := findBook(lib.Books, title)
	if err != nil {
		return err
	}
	member, err := findMember(lib.Members, memberName)
	if err != nil {
		return err
	}
	book.checkedOut = true
	book.checkedOutDate = time.Now()
	book.checkedOutBy = member
	return nil
}

func CheckInBook(lib *Library, title string) error {
	book, err := findBook(lib.Books, title)
	if err != nil {
		return err
	}

	book.checkedOut = false
	book.returnedDate = time.Now()
	book.checkedOutBy = &Member{}

	return nil
}

func findBook(list []*Book, title string) (*Book, error) {
	for _, book := range list {
		if book.title == title {
			return book, nil
		}
	}
	return nil, fmt.Errorf("book not found in library")
}

func findMember(list []*Member, name string) (*Member, error) {
	for _, member := range list {
		if member.name == name {
			return member, nil
		}
	}
	return nil, fmt.Errorf("member not found in library")
}

func main() {
	var lib Library
	AddBook(&lib, "One Fish, Two Fish, Red Fish, Blue Fish", "Dr. Suess")
	AddBook(&lib, "Pachinko", "Min Jin Lee")
	AddBook(&lib, "It", "Stephen King")
	AddBook(&lib, "Post Office", "Charles Bukowski")

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
	err = CheckInBook(&lib, "It")
	if err != nil {
		fmt.Println(err)
	}
	PrintLibrary(&lib)

}
