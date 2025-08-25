package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Simulate CPU-intensive synchronous tasks
func processUserData(id int) map[string]interface{} {
	start := time.Now()
	// Simulate heavy computation
	result := 0.0
	for i := 0; i < 10000000; i++ {
		result += math.Sqrt(float64(i))
	}
	end := time.Now()
	fmt.Printf("processUserData(%d) completed in %v\n", id, end.Sub(start))
	return map[string]interface{}{
		"id":        id,
		"name":      fmt.Sprintf("User %d", id),
		"processed": result,
	}
}

func processUserPosts(id int) map[string]interface{} {
	start := time.Now()
	// Simulate heavy computation
	result := 0.0
	for i := 0; i < 8000000; i++ {
		result += math.Pow(float64(i), 2)
	}
	end := time.Now()
	fmt.Printf("processUserPosts(%d) completed in %v\n", id, end.Sub(start))
	return map[string]interface{}{
		"userId":    id,
		"posts":     []string{fmt.Sprintf("Post 1 by user %d", id), fmt.Sprintf("Post 2 by user %d", id)},
		"processed": result,
	}
}

func processUserSettings(id int) map[string]interface{} {
	start := time.Now()
	// Simulate heavy computation
	result := 0.0
	for i := 0; i < 5000000; i++ {
		result += math.Sin(float64(i))
	}
	end := time.Now()
	fmt.Printf("processUserSettings(%d) completed in %v\n", id, end.Sub(start))
	return map[string]interface{}{
		"userId":        id,
		"theme":         "dark",
		"notifications": true,
		"processed":     result,
	}
}

// Running synchronous tasks in parallel using goroutines and sync.WaitGroup
func processUserDataParallel(userId int) map[string]interface{} {
	fmt.Println("Starting parallel execution...")
	startTime := time.Now()

	var wg sync.WaitGroup
	var userData, userPosts, userSettings map[string]interface{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		userData = processUserData(userId)
	}()

	go func() {
		defer wg.Done()
		userPosts = processUserPosts(userId)
	}()

	go func() {
		defer wg.Done()
		userSettings = processUserSettings(userId)
	}()

	wg.Wait()

	endTime := time.Now()
	fmt.Printf("All parallel tasks completed in %v\n", endTime.Sub(startTime))

	return map[string]interface{}{
		"user":     userData,
		"posts":    userPosts,
		"settings": userSettings,
	}
}

// Alternative: Using channels for parallel execution
func processUserDataParallelWithChannels(userId int) map[string]interface{} {
	fmt.Println("Starting parallel execution with channels...")
	startTime := time.Now()

	userDataChan := make(chan map[string]interface{})
	userPostsChan := make(chan map[string]interface{})
	userSettingsChan := make(chan map[string]interface{})

	go func() {
		userDataChan <- processUserData(userId)
	}()

	go func() {
		userPostsChan <- processUserPosts(userId)
	}()

	go func() {
		userSettingsChan <- processUserSettings(userId)
	}()

	userData := <-userDataChan
	userPosts := <-userPostsChan
	userSettings := <-userSettingsChan

	endTime := time.Now()
	fmt.Printf("All parallel tasks completed in %v\n", endTime.Sub(startTime))

	return map[string]interface{}{
		"user":     userData,
		"posts":    userPosts,
		"settings": userSettings,
	}
}

// Sequential execution (slower)
func processUserDataSequential(userId int) map[string]interface{} {
	fmt.Println("Starting sequential execution...")
	startTime := time.Now()

	userData := processUserData(userId)
	userPosts := processUserPosts(userId)
	userSettings := processUserSettings(userId)

	endTime := time.Now()
	fmt.Printf("Sequential tasks completed in %v\n", endTime.Sub(startTime))

	return map[string]interface{}{
		"user":     userData,
		"posts":    userPosts,
		"settings": userSettings,
	}
}

func main() {
	fmt.Println("=== Go Sync Parallel Example ===")
	
	// Run parallel example with WaitGroup
	result := processUserDataParallel(123)
	fmt.Printf("Parallel result: %+v\n", result)

	fmt.Println("\n=== Go Sync Parallel with Channels Example ===")
	result2 := processUserDataParallelWithChannels(456)
	fmt.Printf("Channels result: %+v\n", result2)

	fmt.Println("\n=== Go Sync Sequential Example ===")
	result3 := processUserDataSequential(789)
	fmt.Printf("Sequential result: %+v\n", result3)
} 