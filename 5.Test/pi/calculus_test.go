package pi

import (
	"math"
	"testing"
)

// Tolerance for comparing floating-point numbers as a percentage
const tolerancePercentage = 0.01 // 1%

// withinTolerance checks if the result is within the allowed percentage tolerance.
func withinTolerance(expected, actual float64) bool {
	margin := tolerancePercentage * expected
	return math.Abs(expected-actual) <= margin
}

func TestRecursivePi(t *testing.T) {
	tests := []struct {
		terms    int
		expected float64
	}{
		{1, 4.0},             // 1 term
		{100, 3.0418396189},  // Approximation with 100 terms
		{1000, 3.1315929036}, // Approximation with 1000 terms
	}

	for _, test := range tests {
		result := RecursivePi(test.terms)
		if !withinTolerance(math.Pi, result) {
			t.Errorf("RecursivePi(%d) = %f; want approximately %f", test.terms, result, math.Pi)
		}
	}
}

func TestIterativePi(t *testing.T) {
	tests := []struct {
		terms    int
		expected float64
	}{
		{1, 4.0},             // 1 term
		{100, 3.0418396189},  // Approximation with 100 terms
		{1000, 3.1315929036}, // Approximation with 1000 terms
	}

	for _, test := range tests {
		result := IterativePi(test.terms)
		if !withinTolerance(math.Pi, result) {
			t.Errorf("IterativePi(%d) = %f; want approximately %f", test.terms, result, math.Pi)
		}
	}
}
