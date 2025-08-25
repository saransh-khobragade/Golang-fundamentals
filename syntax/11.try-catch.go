package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Try-catch (Go uses defer and panic/recover)
func main() {
	reader := bufio.NewReader(os.Stdin)
	
	// Basic error handling
	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number!")
		return
	}
	
	if number == 0 {
		fmt.Println("Cannot divide by zero!")
		return
	}
	
	result := 10.0 / float64(number)
	fmt.Printf("Result: %.2f\n", result)
	
	// This always runs (equivalent to finally)
	defer func() {
		fmt.Println("This always runs!")
	}()
} 