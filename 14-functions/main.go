package main

import "fmt"

var s = 645  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 19 // Game Kills
	gd := 13 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
		// fmt.Println("Manual Sores:", 8783-8142)
	} else if gk == gd {
		fmt.Println("It's not bad but not good", s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}
	if (tb - s) == 0 {
		fmt.Println("CONGRATS YOU ARE CAUGHT UP!!!")
	} else {
		defer fmt.Println("You are ", tb-s, "away from catching up!")
	}
}
