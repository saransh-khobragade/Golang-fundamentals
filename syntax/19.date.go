package main

import (
	"fmt"
	"time"
)

// Date handling
func main() {
	x := time.Now()
	fmt.Println(x)

	y := time.Date(2020, 5, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(y)
} 