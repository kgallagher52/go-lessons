package main

import "fmt"

func main() {
	p1 := struct { // * Creating a anonymous struct
		first string
		last  string
		age   int
	}{
		first: "Ron",
		last:  "Weasley",
		age:   26,
	}
	fmt.Println(p1)
}
