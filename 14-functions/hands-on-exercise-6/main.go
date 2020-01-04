package main

import "fmt"

// 1.Build and use an anonymous func

func main() {
	func() { // Anonymous func
		fmt.Println("Anmonymous func ran")
	}()
}
