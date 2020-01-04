package main

import "fmt"

var m = 11
var d = 03
var y = 1986

func main() {
	fmt.Println(m)
	fmt.Printf("%T\n", m)  // Type
	fmt.Printf("%b\n", m)  // Binary
	fmt.Printf("%x\n", m)  // Hexidecimal
	fmt.Printf("%#x\n", m) // Hexidecimal with x in front
	// fmt.Printf("%T\t%b\t%x\t", m, m, m) // Type - Binary - Hexidecimal
	// fmt.Printf("%x\t%x\t%x\t", m, d, y) // Birthday in hex
	fmt.Printf("%b\t%b\t%b\t", m, d, y) // Birthday in Binary

	// Formating Strings: %T %b %x %#x
	// Println - like print but prints spaces between arguments and a newline at the end
	// Printf - format printing
}
