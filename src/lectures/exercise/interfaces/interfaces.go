//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Vehicle interface {
	MoveToLift()
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) MoveToLift() {
	fmt.Printf("%v is moved to a small lift.\n", m)
}

func (c Car) MoveToLift() {
	fmt.Printf("%v is moved to a standard lift.\n", c)
}

func (t Truck) MoveToLift() {
	fmt.Printf("%v is moved to a large lift.\n", t)
}

func SendAllVehiclesToLifts(lot []Vehicle) {
	fmt.Println("Lifts are empty. Send vehicles to the lifts!")
	for _, v := range lot {
		v.MoveToLift()
	}
}

func main() {
	waitingLine := []Vehicle{Motorcycle("Crotch Rocket"), Car("Turbo Diesel"), Truck("Big Rig")}
	SendAllVehiclesToLifts(waitingLine)
}
