package main

import "fmt"

// if-else and switch
func main() {
	x := 10

	// if-else
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// switch
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is not 1 or 2")
	}

	// switch with expression
	switch {
	case x == 1:
		fmt.Println("x is 1")
	case x == 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is not 1 or 2")
	}
} 