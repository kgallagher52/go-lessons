package main

import "fmt"

var s = 545  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 17 // Game Kills
	gd := 18 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
		// fmt.Println(8247 - 7744)
	} else if gk == gd {
		fmt.Println("It's not bad but not good", s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}

	defer fmt.Println("You are ", tb-s, "away from catching up!")
}
