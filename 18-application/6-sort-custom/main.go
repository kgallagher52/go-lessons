package main

import (
	"fmt"
	"sort"
)

// Bender ...
type Bender struct {
	Name    string
	Element string
	Age     int
}

// ByAge ...
type ByAge []Bender

// How it's sorting it when I call the reciver it's running these to sort by age
func (ba ByAge) Len() int           { return len(ba) }               // Calls this first
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }  // Then this
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age } // Then finally this

// ByName ...
type ByName []Bender

func (bn ByName) Len() int           { return len(bn) }                 // Calls this first
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }    // Then this
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name } // Then finally this

func main() {
	b1 := Bender{Name: "Katara", Element: "Water", Age: 24}
	b2 := Bender{Name: "Sakka", Element: "None", Age: 21}
	b3 := Bender{Name: "Zuko", Element: "Fire", Age: 27}
	b4 := Bender{Name: "Toph", Element: "Earth", Age: 17}

	benders := []Bender{b1, b2, b3, b4}

	// Sort all by age
	fmt.Println("Original: ", benders)
	sort.Sort(ByAge(benders))
	fmt.Println("Sorted by Age: ", benders)

	// Sort all by name
	sort.Sort(ByName(benders))
	fmt.Println("Sorted by Age & Name", benders)

}
