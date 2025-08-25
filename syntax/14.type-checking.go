package main

import (
	"fmt"
	"reflect"
)

// Type checking
func main() {
	a := 1
	b := "2"
	c := []int{1, 2, 3}
	d := map[int]bool{1: true, 2: true, 3: true} // set equivalent
	e := map[string]string{"a": "1", "b": "2", "c": "3"}

	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // string
	fmt.Printf("%T\n", c) // []int
	fmt.Printf("%T\n", d) // map[int]bool
	fmt.Printf("%T\n", e) // map[string]string

	// Type checking
	if reflect.TypeOf(a) == reflect.TypeOf(0) {
		fmt.Println("a is an integer")
	}
	if reflect.TypeOf(b) == reflect.TypeOf("") {
		fmt.Println("b is a string")
	}

	// Alternative type checking
	switch v := a.(type) {
	case int:
		fmt.Println("a is an integer")
	default:
		_ = v
	}
} 