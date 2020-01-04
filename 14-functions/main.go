package main

import "fmt"

var s = 446  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 18 // Game Kills
	gd := 14 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}
}
