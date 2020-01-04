package main

import "fmt"

var s = 399  // Seperation
var tb = 772 // Score to Beat

func main() {
	gk := 8  // Game Kills
	gd := 11 // Game Deaths

	if gk > gd {
		fmt.Println("Good Game!", (gk-gd)+s)
	} else {
		fmt.Println("Bad Game!", s-(gd-gk))
	}
}
