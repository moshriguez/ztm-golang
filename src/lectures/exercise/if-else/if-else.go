//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekday(day int) bool {
	return day < 5
}

func guestDay(day int) bool {
	switch day {
	case 0, 2, 4:
		return true
	default:
		return false
	}
}

func main() {
	// The day and role. Change these to check your work.
	today, role := Saturday, Member
	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		accessGranted()
	} else if role == Member && weekday(today) {
		accessGranted()
	} else if role == Guest && guestDay(today) {
		accessGranted()
	} else {
		accessDenied()
	}
}
