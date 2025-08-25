package main

import "fmt"

// Set operations
func main() {
	a := map[int]bool{1: true, 2: true, 3: true, 4: true}
	b := make(map[int]bool)

	b[4] = true
	b[6] = true

	// Common elements (intersection)
	intersection := make(map[int]bool)
	for key := range a {
		if b[key] {
			intersection[key] = true
		}
	}
	fmt.Println(intersection)

	// Difference
	difference := make(map[int]bool)
	for key := range a {
		if !b[key] {
			difference[key] = true
		}
	}
	fmt.Println(difference)

	s1 := map[int]bool{1: true, 2: true, 3: true} // create set from slice
	fmt.Println(s1)

	s2 := map[int]bool{1: true, 2: true, 3: true} // create set from slice
	fmt.Println(s2)

	s3 := make(map[string]bool) // create set from string
	for _, char := range "ABC" {
		s3[string(char)] = true
	}
	fmt.Println(s3)

	// Converting slice to set
	s4 := map[int]bool{1: true, 2: true, 33: true, 4: true}
	fmt.Println(s4)
} 