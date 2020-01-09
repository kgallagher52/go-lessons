package main

import (
	"fmt"
	"log"
	"strconv"
)

// Stringer ...
type Stringer interface {
	String() string
}

// Book ...
type Book struct {
	Title  string
	Author string
}

// Count ...
type Count int

// Methods
func (b Book) String() string {
	fmt.Println("1. BOOK INVOKED")
	return fmt.Sprintf("2. Book: %s - %s", b.Title, b.Author)
}

func (c Count) String() string {
	fmt.Println("COUNT INVOKED")
	return strconv.Itoa(int(c))
}

// WriteLog ... Functions
func WriteLog(s Stringer) {
	log.Println(s)

	// log.Println(s.String())
}

func main() {
	// Initialize a Count object and pass it to WriteLog().
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	// Initialize a Count object and pass it to WriteLog().
	count := Count(3)
	WriteLog(count)
}
