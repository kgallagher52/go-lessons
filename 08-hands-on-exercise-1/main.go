package main

import "fmt"

func main() {
	// 1. Print every number from 1 to 10,000
	n := 1
	for n <= 10000 {
		fmt.Println("Num:", n)
		n++
	}
}
