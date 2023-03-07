package main

import "fmt"

func main() {
	x := 0

	for x < 100 {
		fmt.Printf("The value of x is %v\n", x)
		x++
	}
}
