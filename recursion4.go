//Printing Ping Pong using Go routine and Recursion in Golang 
package main

import (
	"fmt"
	"time"
)

func print() {
	fmt.Println("Ping\n")
	//delay the go rotine
	time.Sleep(3 * time.Second)
	go fmt.Println("Pong")
	//calling func recursively
	print()

}
func main() {
	print()
}
