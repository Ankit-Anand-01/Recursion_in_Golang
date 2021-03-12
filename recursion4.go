package main

import (
	"fmt"
	"time"
)

func print() {
	fmt.Println("Ping\n")
	time.Sleep(3 * time.Second)
	go fmt.Println("Pong")
	print()

}
func main() {
	print()
}
