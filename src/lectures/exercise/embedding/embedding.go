//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) average() Bytes {
	var sum Bytes = 0
	for _, sample := range b.amount {
		sum += sample
	}
	return sum / Bytes(len(b.amount))
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) average() Celcius {
	var sum Celcius = 0
	for _, sample := range c.temp {
		sum += sample
	}
	return sum / Celcius(len(c.temp))
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) average() Bytes {
	var sum Bytes = 0
	for _, sample := range m.amount {
		sum += sample
	}
	return sum / Bytes(len(m.amount))
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (d *Dashboard) Print() {
	fmt.Println("Average Bandwidth Usage:", d.BandwidthUsage.average())
	fmt.Println("Average CPU Temp:", d.CpuTemp.average())
	fmt.Println("Average Memory Usage:", d.MemoryUsage.average())
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{bandwidth, temp, memory}
	dash.Print()
}
