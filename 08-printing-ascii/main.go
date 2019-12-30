package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\t%#x\n", i, i, i) // Printing out all acsii characters and there code points as well as Hex
	}
	fmt.Println("Hello Moto")
}
