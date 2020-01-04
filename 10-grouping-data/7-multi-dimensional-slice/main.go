package main

import "fmt"

func main() {
	jb := []string{"Harry", "Ron", "Dumbldore", "Snape"}
	fmt.Println(jb)

	mp := []string{"Voldmort", "Bellatrix", "Sirius", "Mallfloy"}
	fmt.Println(mp)

	xp := [][]string{jb, mp} // * Multi-Dimensional String
	fmt.Println(xp)          // Creates a array of two arrays

}
