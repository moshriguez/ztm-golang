//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package testing

import "testing"

func TestGainHealth(t *testing.T) {
	p := Player{
		health: 0,
		maxHealth: 100,
	}
	p.GainHealth(50)
	if p.health != 50 {
		t.Errorf("player's health did not increase correctly. got: %v, expected: %v", p.health, 50)
	}
	p.GainHealth(51)
	if p.health != p.maxHealth {
		t.Errorf("player's health cannot go above the max health. got: %v, expected: %v", p.health, p.maxHealth)
	}
}

func TestLoseHealth(t *testing.T) {
	p := Player{
		health: 100,
		maxHealth: 100,
	}
	p.LoseHealth(50)
	if p.health != 50 {
		t.Errorf("player's health did not decrease correctly. got: %v, expected: %v", p.health, 50)
	}
	p.LoseHealth(51)
	if p.health != 0 {
		t.Errorf("player's health cannot go above the below 0. got: %v", p.health)
	}
}

func TestGainEnergy(t *testing.T) {
	p := Player{
		energy: 0,
		maxEnergy: 100,
	}
	p.GainEnergy(50)
	if p.energy != 50 {
		t.Errorf("player's energy did not increase correctly. got: %v, expected: %v", p.energy, 50)
	}
	p.GainEnergy(51)
	if p.energy != p.maxEnergy {
		t.Errorf("player's energy cannot go above the max energy. got: %v, expected: %v", p.energy, p.maxEnergy)
	}
}

func TestLoseEnergy(t *testing.T) {
	p := Player{
		energy: 100,
		maxEnergy: 100,
	}
	p.LoseEnergy(50)
	if p.energy != 50 {
		t.Errorf("player's energy did not decrease correctly. got: %v, expected: %v", p.energy, 50)
	}
	p.LoseEnergy(51)
	if p.energy != 0 {
		t.Errorf("player's energy cannot go above the below 0. got: %v", p.energy)
	}
}

