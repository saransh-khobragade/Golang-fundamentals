package main

import "fmt"

// Loops
func main() {
	// While loop
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
	// 0
	// 1
	// 2

	// For loop for elements
	fruits := []string{"apple", "banana", "cherry"}
	for _, x := range fruits {
		fmt.Println("first " + x)
	}
	// first apple
	// first banana
	// first cherry

	// For loop for index
	for j := 0; j < len(fruits); j++ {
		fmt.Println(j)
	}
	// 0
	// 1
	// 2

	// For loop for index and element
	for k, v := range fruits {
		fmt.Println(k, v)
	}
	// 0 apple
	// 1 banana
	// 2 cherry

	// For loop for range
	for l := 0; l < 2; l++ {
		fmt.Println(l)
	}
	// 0
	// 1

	// For loop for range
	for m := 2; m < 4; m++ {
		fmt.Println("third", m)
	}
	// third 2
	// third 3

	// For loop for range with step
	for n := 2; n < 6; n += 2 {
		fmt.Println("fourth", n)
	}
	// fourth 2
	// fourth 4

	// For loop in dictionary/object
	d := map[string]string{
		"name": "John",
		"age":  "30",
		"city": "New York",
	}
	for key, value := range d {
		fmt.Println(key, value)
	}
	// name John
	// age 30
	// city New York
} 