package main

import (
	"fmt"
	"strings"
)

// Log
func main() {
	// Hello world
	fmt.Println("Hello world") // Hello world
	fmt.Println("Hello", 23)   // Hello 23

	// Inline string substitution
	fmt.Printf("Hello %d\n", 23) // Hello 23

	// String interpolation
	abc := fmt.Sprintf("Hello my age is %d", 10)
	fmt.Println(abc)

	// Comments
	// This is a single line comment
	fmt.Println("Hello, World!")

	// This is a multiline comment
	/*
		Multiline comment
	*/

	// String concatenation
	name := "John"
	age := 25
	fmt.Println("Hello, " + name + "! You are " + fmt.Sprintf("%d", age) + " years old.")

	// Template strings (using fmt.Sprintf)
	template := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(template)
} 