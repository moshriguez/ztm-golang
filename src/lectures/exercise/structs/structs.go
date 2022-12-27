//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
)

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(rectangle *Rectangle) int {
	return (rectangle.b.x - rectangle.a.x)
}

func length(rectangle *Rectangle) int {
	return (rectangle.a.y - rectangle.b.y)
}


func Area(rectangle *Rectangle) int {
	area := length(rectangle) * width(rectangle)
	fmt.Println("Area of rectangle,", rectangle, ":", area)
	return area
}

func Perimeter(rectangle *Rectangle) int {
	perimeter := (length(rectangle) + width(rectangle)) * 2
	fmt.Println("Perimeter of rectangle,", rectangle, ":", perimeter)
	return perimeter
}

func Double(rectangle *Rectangle) {
	length := length(rectangle)
	width := width(rectangle)
	rectangle.b.x += length
	rectangle.b.y -= width
}

func main() {
	coorA := Coordinate{1, 5}
	coorB := Coordinate{5, 1}
	square := Rectangle{coorA, coorB}
	Area(&square)
	Perimeter(&square)
	Double(&square)
	Area(&square)
	Perimeter(&square)

}
