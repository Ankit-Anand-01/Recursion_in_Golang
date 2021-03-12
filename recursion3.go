//Recursion with base condition in Golang
package main

import "fmt"

func print(n int) {
	//base condition
	if n > 0 {
		fmt.Println("Good Morning! Ankit\n Good Night! Ankit ")
		print(n - 1)
	}
}
func main() {
	print(8)
}
