package main

import "fmt"

func main() {
	//These are all pre-declared constents
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	// ___________________
	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else {
		fmt.Println("our value was not 40")
	}

	y := 434
	if y == 40 {
		fmt.Println("our value was 40")
	} else if y == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40 or 41")
	}

	s := 434
	if s == 40 {
		fmt.Println("our value was 40")
	} else if s == 41 {
		fmt.Println("our value was 41")
	} else if s == 42 {
		fmt.Println("our value was 42")
	} else if s == 43 {
		fmt.Println("our value was 43")
	} else {
		fmt.Println("our value was not 40, 41, 42, or 43")
	}
}
