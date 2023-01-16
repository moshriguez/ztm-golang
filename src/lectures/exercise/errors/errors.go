//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Time struct {
	hour   int
	minute int
	second int
}

func TimeFromString(str string) (Time, error) {
	var time Time
	// validation
	reg, err := regexp.Compile(`\d{2}:\d{2}:\d{2}`)
	if err != nil {
		return Time{}, fmt.Errorf("Trouble with the ol' regexp maker: %v", err)
	}
	if !reg.MatchString(str) {
		return Time{}, errors.New("bro, your string is not in the correct format. Try 'HH:MM:SS'. only use numbers and don't forget those starting zeros!")
	}

	digitalTime := strings.Split(str, ":")
	for i, deet := range digitalTime {
		d, err := strconv.Atoi(deet)
		if err != nil {
			return Time{}, err
		}
		switch i {
		case 0:
			if d > 23 || d < 0 {
				return Time{}, fmt.Errorf("invalid hour")
			}
			time.hour = d
		case 1:
			if d > 59 || d < 0 {
				return Time{}, fmt.Errorf("invalid minute")
			}
			time.minute = d
		case 2:
			if d > 59 || d < 0 {
				return Time{}, fmt.Errorf("invalid second")
			}
			time.second = d
		default:
			return Time{}, fmt.Errorf("Time string should only have 3 parts. Invalid index found: %v", i)
		}
	}
	return time, nil
}
