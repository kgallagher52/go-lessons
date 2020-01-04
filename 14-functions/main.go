package main

import "fmt"

var s = 641  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 33 // Game Kills
	gd := 20 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
		// fmt.Println("Manual Sores:", 8783-8142)
	} else if gk == gd {
		fmt.Println("It's not bad but not good", s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}

	defer fmt.Println("You are ", tb-s, "away from catching up!")
}
