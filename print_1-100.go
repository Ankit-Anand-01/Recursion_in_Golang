//printing 1 to 100 without using loop
package main

import "fmt"

func printing(n int) {
	if n <= 100 {
		fmt.Println(n)
		n++
		printing(n)
	}
}

func main() {
	printing(1)
}
