package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Input handling
func main() {
	reader := bufio.NewReader(os.Stdin)

	// Example 1: Reading two space-separated integers
	fmt.Print("Enter two integers (space-separated): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) >= 2 {
		if p, err := strconv.Atoi(parts[0]); err == nil {
			if c, err := strconv.Atoi(parts[1]); err == nil {
				fmt.Printf("p = %d, c = %d\n", p, c)
			} else {
				fmt.Println("Error: Please enter two space-separated integers")
			}
		} else {
			fmt.Println("Error: Please enter two space-separated integers")
		}
	} else {
		fmt.Println("Error: Please enter two space-separated integers")
	}

	// Example 2: Reading space-separated integers into a slice
	fmt.Print("Enter space-separated integers: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts = strings.Fields(input)
	var num []int
	for _, part := range parts {
		if val, err := strconv.Atoi(part); err == nil {
			num = append(num, val)
		} else {
			fmt.Println("Error: Please enter space-separated integers only")
			return
		}
	}
	fmt.Printf("Slice of integers: %v\n", num)

	// Example 3: Reading a string and converting to slice of characters
	fmt.Print("Enter a string: ")
	exp, _ := reader.ReadString('\n')
	exp = strings.TrimSpace(exp)
	var charList []string
	for _, char := range exp {
		charList = append(charList, string(char))
	}
	fmt.Printf("Slice of characters: %v\n", charList)

	// Example 4: Reading a single integer
	fmt.Print("Enter a single integer: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if singleNum, err := strconv.Atoi(input); err == nil {
		fmt.Printf("Single integer: %d\n", singleNum)
	} else {
		fmt.Println("Error: Please enter a valid integer")
	}
} 