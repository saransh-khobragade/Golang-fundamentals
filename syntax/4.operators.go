package main

import "fmt"

// Operators
func main() {
	// Arithmetic operators
	a := 10
	b := 5
	fmt.Println(a + b) // 15
	fmt.Println(a - b) // 5
	fmt.Println(a * b) // 50
	fmt.Println(float64(a) / float64(b)) // 2.0
	fmt.Println(a % b) // 0 (mod)
	
	// Power (using math.Pow)
	power := 1
	for i := 0; i < b; i++ {
		power *= a
	}
	fmt.Println(power) // 100000

	// Assignment operators
	a = 10
	b = 5
	a += b // a = a + b
	fmt.Println(a) // 15
	a -= b // a = a - b
	fmt.Println(a) // 10
	a *= b // a = a * b
	fmt.Println(a) // 50
	a /= b // a = a / b
	fmt.Println(a) // 10

	// Comparison operators
	a = 10
	b = 5
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false

	// Logical operators
	boolA := true
	boolB := false
	fmt.Println(boolA && boolB) // false
	fmt.Println(boolA || boolB) // true
	fmt.Println(!boolA)         // false

	// Identity operators (Go doesn't have direct equivalent, using comparison)
	numA := 10
	numB := 5
	fmt.Println(numA == numB) // false
	fmt.Println(numA != numB) // true

	// Membership operators (for slices)
	list := []int{1, 2, 3, 4, 5}
	found := false
	for _, val := range list {
		if val == a {
			found = true
			break
		}
	}
	fmt.Println(found) // false
	fmt.Println(!found) // true
} 