package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate async tasks that take different amounts of time
func fetchUserData(id int) map[string]interface{} {
	time.Sleep(2 * time.Second)
	return map[string]interface{}{
		"id":    id,
		"name":  fmt.Sprintf("User %d", id),
		"email": fmt.Sprintf("user%d@example.com", id),
	}
}

func fetchUserPosts(id int) map[string]interface{} {
	time.Sleep(1500 * time.Millisecond)
	return map[string]interface{}{
		"userId": id,
		"posts":  []string{fmt.Sprintf("Post 1 by user %d", id), fmt.Sprintf("Post 2 by user %d", id)},
	}
}

func fetchUserSettings(id int) map[string]interface{} {
	time.Sleep(1 * time.Second)
	return map[string]interface{}{
		"userId":        id,
		"theme":         "dark",
		"notifications": true,
	}
}

// Running tasks in parallel using goroutines and channels
func loadUserDataParallel(userId int) map[string]interface{} {
	fmt.Println("Starting parallel execution...")
	startTime := time.Now()

	// Create channels to receive results
	userDataChan := make(chan map[string]interface{})
	userPostsChan := make(chan map[string]interface{})
	userSettingsChan := make(chan map[string]interface{})

	// Launch goroutines for parallel execution
	go func() {
		userDataChan <- fetchUserData(userId)
	}()

	go func() {
		userPostsChan <- fetchUserPosts(userId)
	}()

	go func() {
		userSettingsChan <- fetchUserSettings(userId)
	}()

	// Wait for all results
	userData := <-userDataChan
	userPosts := <-userPostsChan
	userSettings := <-userSettingsChan

	endTime := time.Now()
	fmt.Printf("All tasks completed in %v\n", endTime.Sub(startTime))

	return map[string]interface{}{
		"user":     userData,
		"posts":    userPosts,
		"settings": userSettings,
	}
}

// Alternative: Using sync.WaitGroup for parallel execution
func loadUserDataParallelWithWaitGroup(userId int) map[string]interface{} {
	fmt.Println("Starting parallel execution with WaitGroup...")
	startTime := time.Now()

	var wg sync.WaitGroup
	var userData, userPosts, userSettings map[string]interface{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		userData = fetchUserData(userId)
	}()

	go func() {
		defer wg.Done()
		userPosts = fetchUserPosts(userId)
	}()

	go func() {
		defer wg.Done()
		userSettings = fetchUserSettings(userId)
	}()

	wg.Wait()

	endTime := time.Now()
	fmt.Printf("All tasks completed in %v\n", endTime.Sub(startTime))

	return map[string]interface{}{
		"user":     userData,
		"posts":    userPosts,
		"settings": userSettings,
	}
}

// Sequential execution (slower)
func loadUserDataSequential(userId int) map[string]interface{} {
	fmt.Println("Starting sequential execution...")
	startTime := time.Now()

	userData := fetchUserData(userId)
	userPosts := fetchUserPosts(userId)
	userSettings := fetchUserSettings(userId)

	endTime := time.Now()
	fmt.Printf("Sequential tasks completed in %v\n", endTime.Sub(startTime))

	return map[string]interface{}{
		"user":     userData,
		"posts":    userPosts,
		"settings": userSettings,
	}
}

func main() {
	fmt.Println("=== Go Async Parallel Example ===")
	
	// Run parallel example
	result := loadUserDataParallel(123)
	fmt.Printf("Parallel result: %+v\n", result)

	fmt.Println("\n=== Go Async Parallel with WaitGroup Example ===")
	result2 := loadUserDataParallelWithWaitGroup(456)
	fmt.Printf("WaitGroup result: %+v\n", result2)

	fmt.Println("\n=== Go Sequential Example ===")
	result3 := loadUserDataSequential(789)
	fmt.Printf("Sequential result: %+v\n", result3)
} 