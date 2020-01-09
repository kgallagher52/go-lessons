package main

import "fmt"

func main() {
	x := 0       // 1.1 Assign x too 0
	firstStep(x) // 1.2 Invoke firstStep and pass x
	fmt.Println("1.3", x)

	z := 0         // 2.1 Assign z too 0
	secondStep(&z) // 2.2 Invoke secondStep and pass z which is now an address to z
	fmt.Println("2.3", z)

}

// 1.3 set y to the value of x passed in
func firstStep(y int) {
	fmt.Println("1.1", y)
	y = 52 // 1.4 assign y to 52
	fmt.Println("1.2", y)
}

// 2.3 set p to the value of z passed in
func secondStep(p *int) {
	fmt.Println("2.1", p)
	*p = 52 // 2.4 assign p to 52
	fmt.Println("2.2", *p)
}
