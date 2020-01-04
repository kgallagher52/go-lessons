package main

import "fmt"

var s = 484  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 19 // Game Kills
	gd := 10 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}
}
