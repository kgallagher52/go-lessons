package main

import "fmt"

func main() {
	s := "Hello World!"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // Converting s into bytes

	fmt.Println(bs)
	fmt.Printf("%T\n", bs) // uint8 - which is an alias for UTF-8 coding schema

	for i := 0; i < len(s); i++ { // Printing out the UTF-8 Code points
		fmt.Printf("%#U\n", s[i])
	}

	for i, v := range s { // SHowing the hex value for each letter
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
