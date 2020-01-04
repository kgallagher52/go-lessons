package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ { // Basic loop all the way up to 100
		// fmt.Println(i)
	}

	// Nesting loops
	for i := 0; i <= 10; i++ {
		// fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j <= 3; j++ {
			// fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	// Single condition Slight way of doing a while loop in Golang
	x := 1
	for x < 10 {
		// fmt.Println("X", x)
		x++
	}
	// fmt.Println("Done")

	// Another way of writing this single conditon
	y := 1
	for {
		if y > 10 {
			break // Break outside the code
		}
		// fmt.Println("Y", y)
		y++
	}
	// fmt.Println("done")

	// Using continue keyword
	z := 1
	for {
		z++
		if z > 100 {
			break
		}
		if z%2 != 0 {
			continue // If this is true it will not go to the next code block rather it will move back to the top and start the proccess over again
		}
		fmt.Printf("%d\t", z)
	}
	fmt.Println("Z Done.")
}
