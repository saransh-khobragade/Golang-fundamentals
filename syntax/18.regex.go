package main

import (
	"fmt"
	"regexp"
)

// Regular expressions
func main() {
	// Return if specific word present
	txt := "The rain in Spain Spainish 01 "
	pattern1 := regexp.MustCompile(`\bSpai\w+`)
	matches1 := pattern1.FindString(txt)
	if matches1 != "" {
		fmt.Println(matches1)
		loc := pattern1.FindStringIndex(txt)
		fmt.Printf("%d %d\n", loc[0], loc[1])
	}
	fmt.Println("------------")

	// Return if digits
	pattern2 := regexp.MustCompile(`\d+`)
	matches2 := pattern2.FindString(txt)
	if matches2 != "" {
		fmt.Println(matches2)
		loc := pattern2.FindStringIndex(txt)
		fmt.Printf("%d %d\n", loc[0], loc[1])
	}
	fmt.Println("------------")

	// Return all occurrences
	txt2 := "The rain in Spain"
	pattern3 := regexp.MustCompile("ai")
	matches3 := pattern3.FindAllString(txt2, -1)
	for _, match := range matches3 {
		fmt.Println(match)
	}
	fmt.Println("------------")

	// Regex patterns:
	// \b - Returns a match where the specified characters are at the beginning or at the end of a word
	// \d - Returns a match where the string contains digits (numbers from 0-9)
	// \s - Returns a match where the string contains a white space character
	// \w - Returns a match where the string contains any word characters (characters from a to Z, digits from 0-9, and the underscore _ character)
} 