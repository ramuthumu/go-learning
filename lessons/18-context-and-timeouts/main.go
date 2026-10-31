package main

import (
	"context"
	"fmt"
	"time"
)

// simulateSlowQuery simulates a database or network query that takes 'workTime' to finish.
// It respects cancellation via ctx.Done().
func simulateSlowQuery(ctx context.Context, queryID string, workTime time.Duration) (string, error) {
	fmt.Printf("[%s] Query started (needs %v)...\n", queryID, workTime)

	select {
	case <-time.After(workTime):
		// Finished in time!
		return fmt.Sprintf("[%s] Query SUCCESS in %v", queryID, workTime), nil
	case <-ctx.Done():
		// Context timed out or was cancelled!
		return "", fmt.Errorf("[%s] Query ABORTED: %w", queryID, ctx.Err())
	}
}

func main() {
	// 1. Context with Timeout:
	// "This entire operation must complete within 150 milliseconds, or cancel!"
	timeoutDuration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel() // Always defer cancel to release context resources!

	fmt.Printf("--- 1. Testing fast operation within %v timeout ---\n", timeoutDuration)
	// Fast query (takes 50ms) -> should succeed
	result, err := simulateSlowQuery(ctx, "FastQuery", 50*time.Millisecond)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 2. Testing slow operation that exceeds timeout
	fmt.Printf("\n--- 2. Testing slow operation that exceeds timeout ---\n")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancelTimeout()

	// Slow query (takes 300ms, but timeout is 100ms) -> should abort!
	_, err = simulateSlowQuery(ctxTimeout, "SlowQuery", 300*time.Millisecond)
	if err != nil {
		fmt.Printf("Expected error caught: %v\n", err)
	}

	// 3. Manual Cancellation (WithCancel)
	fmt.Printf("\n--- 3. Testing manual cancellation ---\n")
	ctxManual, cancelManual := context.WithCancel(context.Background())

	go func() {
		time.Sleep(30 * time.Millisecond)
		fmt.Println("User clicked 'Cancel' button! Calling cancelManual()...")
		cancelManual()
	}()

	_, err = simulateSlowQuery(ctxManual, "UserQuery", 500*time.Millisecond)
	if err != nil {
		fmt.Printf("Query stopped gracefully: %v\n", err)
	}
}
