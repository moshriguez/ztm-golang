//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(item *Item) {
	item.tag = Active
}

func deactivate(item *Item) {
	item.tag = Inactive
}

func checkout(list []*Item) {
	fmt.Println("checking out...")
	for _, item := range list {
		deactivate(item)
	}
}

func addInventory(itemName string) Item {
	i := Item{name: itemName}
	activate(&i)
	return i
}

func printShoppingList(list []*Item) {
	fmt.Println("Shopping List:")
	for _, item := range list {
		fmt.Println(item)
	}
}

func main() {
	catFood := addInventory("cat food")
	bowl := addInventory("bowl")
	waterFountain := addInventory("water fountain")
	kibble := addInventory("kibble")

	shoppingList := []*Item{&catFood, &bowl, &waterFountain, &kibble}
	printShoppingList(shoppingList)

	deactivate(shoppingList[3])
	printShoppingList(shoppingList)

	checkout(shoppingList)
	printShoppingList(shoppingList)
}
