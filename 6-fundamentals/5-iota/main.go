package main

// *iota is used if you need to easily incrment by something starting at 1
import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

// You can also do this for ease of proggraming
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
