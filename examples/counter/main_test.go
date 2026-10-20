package main

import (
	"sync"
	"testing"
)

func TestConcurrentCounter(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Add() }()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Fatalf("got %d, want 1000", got)
	}
}
