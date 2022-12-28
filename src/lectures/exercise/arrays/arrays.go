//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import (
	"fmt"
	"math"
)

type Product struct {
	name string
	price float64
}

func PrintListDetails(list [4]Product) {
	var totalCost float64
	var totalItems int
	for i:=0; i < len(list); i++ {
		item := list[i]
		totalCost += item.price
		if item.name != "" {
			totalItems += 1
		}
	}
	fmt.Println("This list has", totalItems, "items")
	fmt.Println("The last item on the list:", list[totalItems-1].name)
	fmt.Println("Total cost of list:", math.Round(totalCost*100)/100)
}

func main() {
	var shoppingList [4]Product
	shoppingList[0] = Product{"chocolate", 3.50}
	shoppingList[1] = Product{"strowberries", 5}
	shoppingList[2] = Product{"marshmallows", 4.97}

	PrintListDetails(shoppingList)

	shoppingList[3] = Product{"mango", 1.99}

	PrintListDetails(shoppingList)
}
