package main

import "fmt"

// This way of declairation doesn't work outside the function
// myName := "Wonderboy"

func main() {
	// String Declairation
	var name string = "Wonderboy"
	myName := "Wonderboy"
	var stringPlaceholder string

	fmt.Println("My name is", name)
	fmt.Println("My name is", myName)
	fmt.Println("My name is", stringPlaceholder)
	stringPlaceholder = "Wonderboy"
	fmt.Println("My name is", stringPlaceholder)
	fmt.Println("")

	// Integer Declairation

	var age int = 25
	myAge := 25

	var intPlaceholder int

	fmt.Println("My age is ", age)
	fmt.Println("My age is ", myAge)
	fmt.Println("My age is ", intPlaceholder)
	intPlaceholder = 25
	fmt.Println("My age is ", intPlaceholder)
	fmt.Println("")

	// Arrays

	ary := []any{"dog", "cat", "snake", 1, 50}

	fmt.Println(ary)
	fmt.Println(len(ary))
	fmt.Println("")
}
