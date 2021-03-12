//infinite recursive func in Golang
package main

import (
	"fmt"
)

// infinite recursion function
func print() {

	// printing infinite times
	fmt.Println("Ankit")
	//calling func recursively
	print()
}

func main() {
       //calling func
	print()
}
