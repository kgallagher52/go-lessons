package main

import "fmt"

var s = 483  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 13 // Game Kills
	gd := 17 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
		// fmt.Println(8197 - 7710)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}
}
