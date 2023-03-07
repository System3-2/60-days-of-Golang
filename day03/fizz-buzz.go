package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizz buzz", i)
		} else if i % 3 == 0 {
			fmt.Println("buzz", i)
		} else if i % 5 == 0 {
			fmt.Println("fizz", i)
		}

	}
}
