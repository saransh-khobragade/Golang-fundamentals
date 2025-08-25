package main

import (
	"fmt"
	"math"
)

// Variables
func main() {
	integer := 12
	str := "Saransh"
	floatVar := 20.5
	bool := true
	var nullValue interface{} = nil

	// Multiple declaration
	a, b, c := "Orange", "Banana", "Cherry"

	// Arrays/Slices
	arr := []string{"apple", "banana", "cherry"}

	// Tuple equivalent (using interface slice)
	tuple := []interface{}{"apple", "banana", "cherry"}

	// Map (HashMap equivalent)
	hashMap := map[string]interface{}{
		"name": "John",
		"age":  36,
	}

	// Set (using map with bool values)
	set := map[string]bool{
		"apple":  true,
		"banana": true,
		"cherry": true,
	}

	// Infinity
	infinite := math.Inf(1)

	fmt.Println(integer) // 12
	fmt.Println(str)     // Saransh
	fmt.Println(floatVar) // 20.5
	fmt.Println(nullValue) // <nil>
	fmt.Println(a, b, c)   // Orange Banana Cherry
	fmt.Println(arr)       // [apple banana cherry]
	fmt.Println(tuple)     // [apple banana cherry]
	fmt.Println(hashMap)   // map[age:36 name:John]
	fmt.Println(set)       // map[apple:true banana:true cherry:true]
	fmt.Println(infinite)  // +Inf
} 