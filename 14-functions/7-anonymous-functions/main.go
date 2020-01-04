package main

import "fmt"

func main() {
	foo()

	func() { // Anonymous func
		fmt.Println("Anmonymous func ran")
	}()

	func(x int) { // Anonymous function with arguments and parameters
		fmt.Println("The meaning of life:", x)
	}(52)
}

func foo() { // Not an anonymous func because it's named foo
	fmt.Println("foo ran")
}
