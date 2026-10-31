package main

import (
	"fmt"
	"time"
)

// A function that simulates checking a server status
func checkService(name string, delay time.Duration, results chan string) {
	time.Sleep(delay)
	// Send the result message into the channel using <-
	results <- fmt.Sprintf("%s is UP (checked in %v)", name, delay)
}

func main() {
	// 1. Create a channel to communicate between goroutines
	results := make(chan string)

	fmt.Println("Checking services concurrently...")

	// 2. Launch 3 lightweight background tasks using the 'go' keyword!
	go checkService("Plex Media Server", 100*time.Millisecond, results)
	go checkService("Home Assistant", 50*time.Millisecond, results)
	go checkService("Pi-hole DNS", 80*time.Millisecond, results)

	// 3. Receive all 3 messages from the channel as they finish
	for i := 0; i < 3; i++ {
		msg := <-results // Waits for the next result from the channel
		fmt.Println(" ->", msg)
	}

	fmt.Println("All checks complete!")
}
