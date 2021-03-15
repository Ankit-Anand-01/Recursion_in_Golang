//indirect recursion in Golang
package main

import (
	"fmt"
	"time"
)

func printping() {
	fmt.Println("Ping")
	time.Sleep(2 * time.Second)
	printpong()
}
func printpong() {
	fmt.Println("Pong")
	time.Sleep(2 * time.Second)
	printping()
}
func main() {
	printping()
}
