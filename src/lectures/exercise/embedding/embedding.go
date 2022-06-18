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

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}


func (band *BandwidthUsage) GetAverageBandwidth() int {
	var sum int
	for item := range band.amount{
		// Note: not put item! it just counts each iteration!
		// sum+=item
		sum+=int(band.amount[item])
	}
	if len(band.amount) > 0 {
		return sum/len(band.amount)
	}
	return 0
}

func (temp *CpuTemp) GetAverageTemp() int {
	var sum int
	for item := range temp.temp{
		//fmt.Println(temp.temp[item])
		sum+=int(temp.temp[item])
	}
	if len(temp.temp) > 0 {
		return sum/len(temp.temp)
	}
	return 0
}

func (mem *MemoryUsage) GetAverageMemory() int {
	var sum int
	for item := range mem.amount{
		sum+=int(mem.amount[item])
	}
	if len(mem.amount) > 0 {
		return sum/len(mem.amount)
	}
	return 0
}

//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
type Dashboard struct{
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}

	//* Print out a 5-second average from each component using promoted
	//  methods on the Dashboard
	fmt.Printf("Average bandwidth usage: %v\n", dash.GetAverageBandwidth())
	fmt.Printf("Average temp: %v\n", dash.GetAverageTemp())
	fmt.Printf("Average memory usage: %v\n", dash.GetAverageMemory())

	// Note: not possible/or so easy if the methods all had the same name!
}
