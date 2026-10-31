package main

import "testing"

// TestCalculateDiscount demonstrates standard Go "Table-Driven Testing".
// This is the universal Go pattern for testing functions with multiple inputs.
func TestCalculateDiscount(t *testing.T) {
	// Define the test case struct
	tests := []struct {
		name      string
		price     float64
		discount  int
		expected  float64
		expectErr bool
	}{
		{name: "standard 20% discount", price: 100.0, discount: 20, expected: 80.0, expectErr: false},
		{name: "zero discount", price: 50.0, discount: 0, expected: 50.0, expectErr: false},
		{name: "full 100% discount", price: 75.0, discount: 100, expected: 0.0, expectErr: false},
		{name: "invalid negative price", price: -10.0, discount: 10, expected: 0.0, expectErr: true},
		{name: "invalid negative discount", price: 100.0, discount: -5, expected: 0.0, expectErr: true},
		{name: "invalid discount over 100", price: 100.0, discount: 150, expected: 0.0, expectErr: true},
	}

	for _, tc := range tests {
		// t.Run creates a subtest for each table entry
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalculateDiscount(tc.price, tc.discount)

			if tc.expectErr {
				if err == nil {
					t.Fatalf("expected error for case %q, but got nil", tc.name)
				}
				return // Expected error occurred, test passed
			}

			if err != nil {
				t.Fatalf("unexpected error for case %q: %v", tc.name, err)
			}

			if got != tc.expected {
				t.Errorf("case %q failed: got %.2f, expected %.2f", tc.name, got, tc.expected)
			}
		})
	}
}

// BenchmarkCalculateDiscount demonstrates Go's built-in benchmark runner.
// Run with: go test -bench=. ./lessons/17-testing
func BenchmarkCalculateDiscount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = CalculateDiscount(199.99, 15)
	}
}
