package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Randomization
func main() {
	rand.Seed(time.Now().UnixNano())
	
	list1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(list1[rand.Intn(len(list1))]) // random number from slice

	fmt.Println(rand.Intn(9) + 1) // random number from 1 to 10
} 