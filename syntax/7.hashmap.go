package main

import (
	"fmt"
	"sort"
)

// HashMap
func main() {
	// Blank dictionary
	dic := make(map[string]int)
	dic["a"] = 5
	dic["b"] = 10

	// How to loop over dictionary
	for key, value := range dic {
		fmt.Println(key, value)
	}

	for key := range dic {
		fmt.Println(key)
	}

	for _, value := range dic {
		fmt.Println(value)
	}

	// How to check key exists in dictionary
	if _, exists := dic["x"]; exists {
		dic["x"]++
	} else {
		dic["x"] = 1
	}

	// Set unique but no unordered
	setItems := map[string]map[string]bool{
		"A": {"B": true, "C": true},
		"B": {"A": true, "D": true, "E": true},
		"C": {"A": true, "F": true},
		"D": {"B": true},
		"E": {"B": true, "F": true},
		"F": {"C": true, "E": true},
	}
	fmt.Println(setItems)

	// Dictionary key value pairs can use immutable keys mutable values
	dictionary := map[string]int{
		"A": 12,
		"B": 122,
		"C": 45,
		"D": 76,
		"E": 23,
		"F": 2323,
	}
	fmt.Println(dictionary)

	// Get keys and values
	keys := make([]string, 0, len(dictionary))
	values := make([]int, 0, len(dictionary))
	for key, value := range dictionary {
		keys = append(keys, key)
		values = append(values, value)
	}
	fmt.Println(keys)
	fmt.Println(values)

	// Add item or update to dictionary
	dictionary["G"] = 31
	dictionary["G"] = 32

	// Loop in dict
	for key, value := range dictionary {
		fmt.Println(key, value)
	}

	// Sorting dictionary by its values and converting to list of tuples
	type kv struct {
		Key   string
		Value int
	}
	var sortedByValue []kv
	for key, value := range dictionary {
		sortedByValue = append(sortedByValue, kv{key, value})
	}
	sort.Slice(sortedByValue, func(i, j int) bool {
		return sortedByValue[i].Value < sortedByValue[j].Value
	})
	fmt.Println(sortedByValue)

	// Sorting dictionary by its keys and converting to list of tuples
	var sortedByKey []kv
	for key, value := range dictionary {
		sortedByKey = append(sortedByKey, kv{key, value})
	}
	sort.Slice(sortedByKey, func(i, j int) bool {
		return sortedByKey[i].Key < sortedByKey[j].Key
	})
	fmt.Println(sortedByKey)

	// Complex key in dictionary using tuple (using slice as key)
	d := make(map[string]int)
	d["a,b,c"] = 1 // Using string representation of tuple
	fmt.Println(d)
} 