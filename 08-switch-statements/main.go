package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print")
	case (3 == 3):
		fmt.Println("Prints")
		fallthrough // This allows the switch statment to continue. Because if I didn't have that it would print "Prints" and break out of the program since that case was true fallthrough will tell it to keep going
	case (4 == 4):
		fmt.Println("Only Print if there is a fallthrough on the true case above")
		fallthrough
	case (7 == 9):
		fmt.Println("Not true")
		fallthrough
	case (11 == 14):
		fmt.Println("Not true")
		fallthrough
	case (15 == 15):
		fmt.Println("15 True")
	default:
		fmt.Println("This is default")
	}
	//Evaluating on a value
	switch "Harry" {
	case "Tom":
		fmt.Println("This should not print")
	case "Harry":
		fmt.Println("Potter")
	default:
		fmt.Println("This is default Two")
	}

	//Having Multiple cases
	v := "Voldmort"
	switch v {
	case "Tom", "Voldmort", "Bart":
		fmt.Println("Has returned!!!")
	case "Harry":
		fmt.Println("Potter")
	default:
		fmt.Println("This is default Two")
	}

}
