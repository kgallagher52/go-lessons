package main

import "fmt"

var s = 634  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 15 // Game Kills
	gd := 19 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
		// fmt.Println("Manual Sores:", 8633-8023)
	} else if gk == gd {
		fmt.Println("It's not bad but not good", s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}

	defer fmt.Println("You are ", tb-s, "away from catching up!")
}
