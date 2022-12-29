//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printParts(title string, assLine []Part) {
	fmt.Println("\n---", title, "---")
	for i := 0; i < len(assLine); i++ {
		part := assLine[i]
		fmt.Println(part)
	}
}

func main() {
	assemblyLine := []Part{"snorkel", "polyurethane bag", "body lotion"}
	printParts("OG list", assemblyLine)
	assemblyLine = append(assemblyLine, "5 oz can of cat food", "allen wrench")
	printParts("updated list", assemblyLine)
	assemblyLine = assemblyLine[3:]
	printParts("sliced list", assemblyLine)

}
