//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"


type Player struct {
	name      string
	health    int
	maxHealth int
	energy    int
	maxEnergy int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (p *Player) LoseHealth(amount int) {
	// assuming health cannot by negative
	p.health = Max(0, p.health - amount)
	fmt.Println(p.name, "lost", amount, "health")
	fmt.Println(p.name, "now has", p.health, "health")
}

func (p *Player) GainHealth(amount int) {
	// can't have more than the max health
	p.health = Min(p.maxHealth, p.health + amount)
	fmt.Println(p.name, "gained", amount, "health")
	fmt.Println(p.name, "now has", p.health, "health")
}

func (p *Player) LoseEnergy(amount int) {
	// assuming energy cannot by negative
	p.energy = Max(0, p.energy - amount)
	fmt.Println(p.name, "lost", amount, "energy")
	fmt.Println(p.name, "now has", p.energy, "energy")
}

func (p *Player) GainEnergy(amount int) {
	// can't have more than the max energy
	p.energy = Min(p.maxEnergy, p.energy + amount)
	fmt.Println(p.name, "gained", amount, "energy")
	fmt.Println(p.name, "now has", p.energy, "enery")
}

func main() {
	player := Player{
		name: "Sam",
		health: 20,
		maxHealth: 20,
		energy: 10,
		maxEnergy: 10,
	}
	player.LoseHealth(5)
	player.LoseEnergy(4)
	player.GainEnergy(3)
	player.LoseHealth(13)
	player.GainEnergy(5)
	fmt.Println()
	fmt.Println("checking that limits are enforced")
	player.LoseEnergy(16)
	player.LoseHealth(20)
	player.GainEnergy(100)
	player.GainHealth(100)
}
