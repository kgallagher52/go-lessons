package main

import "fmt"

var s = 577  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 14 // Game Kills
	gd := 15 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
		fmt.Println(8469 - 7902)
	} else if gk == gd {
		fmt.Println("It's not bad but not good", s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}

	defer fmt.Println("You are ", tb-s, "away from catching up!")
}
