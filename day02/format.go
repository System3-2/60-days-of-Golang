package main

import "fmt"

func main() {
	phrase := "Hello golang, nice to work with you"
	name := "oloja"
	age := 44

	fmt.Println(phrase)
	fmt.Printf("Hello, my name is %v and i am %v years old", name, age)
	fmt.Printf("Hello, my name is %v and i am a type %T and i am %v years old and a type %T \n", name, name, age, age)
}
