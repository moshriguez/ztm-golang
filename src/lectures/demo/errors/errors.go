package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(i int) (int, error) {
	if i > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at index %v", i))
	} else {
		return s.values[i], nil
	}
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}
