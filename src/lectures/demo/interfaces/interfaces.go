package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Preparing Dishes...")
	for _, dish := range dishes {
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}
func main() {
	dishes := []Preparer{Chicken("chicken wings"), Salad("Cobb Salad")}
	PrepareDishes(dishes)
}
