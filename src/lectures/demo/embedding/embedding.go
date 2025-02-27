package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type Beltsize int
type Shipping int

func (b Beltsize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() Beltsize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam Mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() Beltsize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v Conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)

	// waste := ToxicWaste{3000}
	// automate(&waste)
}
