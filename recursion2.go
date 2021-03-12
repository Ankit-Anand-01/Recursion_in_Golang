package main

import (
	"fmt"
)

// infinite recursion function
func print() {

	// printing infinite times
	fmt.Println("Ankit")
	print()
}

// main function
func main() {

	// call to infinite
	// recursive function
	print()
}
