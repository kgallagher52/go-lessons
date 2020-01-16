package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "herro")
	io.WriteString(os.Stdout, "herro")
}
