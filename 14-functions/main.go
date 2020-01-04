package main

import "fmt"

var s = 602  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 30 // Game Kills
	gd := 10 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
		fmt.Println("Manual Sores:", 8569-7967)
	} else if gk == gd {
		fmt.Println("It's not bad but not good", s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}

	defer fmt.Println("You are ", tb-s, "away from catching up!")
}
