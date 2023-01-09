package display

import "fmt"

func Display(msg string) {
	fmt.Println(msg)
}

// this will not be availabe outside this package because it is named with a small first letter
func hello(msg string) {
	fmt.Println(msg)
}