package main

import (
	"fmt"
	"sort"
)

type wizard struct {
	First  string
	Last   string
	Age    int
	Powers []string
}

// ByAge ...
type ByAge []wizard

func (ba ByAge) Len() int           { return len(ba) }
func (ba ByAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba ByAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

// ByLast ...
type ByLast []wizard

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

func main() {
	w1 := wizard{
		First: "Harry",
		Last:  "Potter",
		Age:   21,
		Powers: []string{
			"Blade Shooting",
			"Dragon Summon",
			"Invisiblity",
		},
	}

	w2 := wizard{
		First: "Hermine",
		Last:  "Granger",
		Age:   22,
		Powers: []string{
			"Repair Anything",
			"Extract Information",
			"Erase Memories",
		},
	}

	w3 := wizard{
		First: "Ron",
		Last:  "Weasly",
		Age:   20,
		Powers: []string{
			"Light Fires",
			"Win Any Game",
			"Breath Underwater",
		},
	}

	wizards := []wizard{w1, w2, w3}

	sort.Sort(ByAge(wizards))
	fmt.Println("Age:")
	for _, w := range wizards {
		fmt.Printf("\t%v %v %v\n", w.Age, w.First, w.Last)
	}

	sort.Sort(ByLast(wizards))
	fmt.Println("Age & Last:")
	for _, w := range wizards {
		fmt.Printf("\t%v %v\n", w.First, w.Last)
		sort.Strings(w.Powers)
		for _, p := range w.Powers {
			fmt.Printf("\t\t%v\n", p)
		}
	}
}
