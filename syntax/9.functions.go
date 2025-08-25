package main

import "fmt"

// Functions

// Function with parameter
func myFunction(fname string) {
	fmt.Println(fname + " Refsnes")
}

// Function with multiple parameters
func myFunctionWithParams(fname, lname string) {
	fmt.Println(fname + " " + lname)
}

// Function with default parameter (Go doesn't have default parameters, using multiple functions)
func myFunctionWithDefault(fname string) {
	myFunctionWithParams(fname, "Refsnes")
}

func main() {
	// Test the functions
	myFunction("John")
	myFunctionWithParams("John", "Doe")
	myFunctionWithDefault("Jane")
} 