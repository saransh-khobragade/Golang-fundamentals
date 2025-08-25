package main

import (
	"fmt"
	"strings"
)

// Strings
func main() {
	// String interpolation
	abc := fmt.Sprintf("Hello %d", 10)
	fmt.Println(abc)

	// Multi line string
	multiLineString := `
This is a multi line string
This is a multi line string`
	fmt.Println(multiLineString)

	// LENGTH OF STRING
	a := "Hello, World!"
	fmt.Println(len(a)) // 13

	// CHECK IN STRING
	txt := "The best things in life are free!"
	if strings.Contains(txt, "free") {
		fmt.Println("Yes, 'free' is present.")
	}
	if !strings.Contains(txt, "abc") {
		fmt.Println("Yes, 'abc' not is present.")
	}

	// Splitting a string into array
	chars := strings.Split(txt, "")
	fmt.Println(chars)
	// [T h e   b e s t   t h i n g s   i n   l i f e   a r e   f r e e !]

	// Joining a list into a string
	arr := []string{"Hello", "World"}
	fmt.Println(strings.Join(arr, " ")) // Hello World

	// Slicing a string
	str := "Saransh"
	fmt.Println(string(str[1])) // a from index 1 only one element
	fmt.Println(string(str[len(str)-2])) // s from last second index only one element
	fmt.Println(str[2:5]) // ran [including:excluding]

	// Trimming a string
	str = "   Hello, World!   "
	fmt.Println(strings.TrimSpace(str)) // "Hello, World!" (removes leading and trailing spaces)
	fmt.Println(strings.TrimLeft(str, " ")) // "Hello, World!   " (removes leading spaces)
	fmt.Println(strings.TrimRight(str, " ")) // "   Hello, World!" (removes trailing spaces)
} 