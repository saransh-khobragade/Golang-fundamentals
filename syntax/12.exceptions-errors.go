package main

import (
	"errors"
	"fmt"
	"strings"
)

// Creating custom errors and throwing errors

// Custom error types
type AgeError struct {
	message string
}

func (e AgeError) Error() string {
	return e.message
}

type InvalidEmailError struct {
	email   string
	message string
}

func (e InvalidEmailError) Error() string {
	return e.message
}

// Throwing built-in errors
func divideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero!")
	}
	return a / b, nil
}

// Throwing custom errors
func checkAge(age int) (int, error) {
	if age < 0 {
		return 0, AgeError{"Age cannot be negative!"}
	}
	if age > 150 {
		return 0, AgeError{"Age seems unrealistic!"}
	}
	return age, nil
}

func validateEmail(email string) (string, error) {
	if !strings.Contains(email, "@") {
		return "", InvalidEmailError{
			email:   email,
			message: fmt.Sprintf("Invalid email format: %s", email),
		}
	}
	return email, nil
}

func main() {
	// Example usage
	if _, err := checkAge(-5); err != nil {
		if ageErr, ok := err.(AgeError); ok {
			fmt.Println("Age error:", ageErr.Error())
		}
	}

	if _, err := validateEmail("invalid-email"); err != nil {
		if emailErr, ok := err.(InvalidEmailError); ok {
			fmt.Println("Email error:", emailErr.Error())
		}
	}

	if _, err := divideNumbers(10, 0); err != nil {
		fmt.Println("Division error:", err.Error())
	}
} 