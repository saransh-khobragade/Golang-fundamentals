package main

import (
	"fmt"
	"strconv"
)

// Type casting
func main() {
	a, _ := strconv.Atoi("1")
	b, _ := strconv.ParseFloat("4.2", 64)
	c := strconv.Itoa(2)
	
	var d []int
	for i := 1; i < 10; i++ {
		d = append(d, i)
	}
	
	e := make([]interface{}, len(d)) // tuple equivalent
	for i, v := range d {
		e[i] = v
	}
	
	f := make(map[int]bool) // set equivalent
	for _, v := range d {
		f[v] = true
	}
	
	g := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	fmt.Println(a, b, c, d, e, f, g)
} 