package pi

import (
	"math"
)

// RecursivePi calculates pi using the Gregory-Leibniz series recursively.
func RecursivePi(n int) float64 {
	if n == 0 || n == 1 {
		return 4.0
	}
	term := 4.0 * math.Pow(-1, float64(n)) / float64(2*n+1)
	return term + RecursivePi(n-1)
}

// IterativePi calculates pi using the Gregory-Leibniz series iteratively.
func IterativePi(terms int) float64 {
	pi := 0.0

	for i := 0; i < terms; i++ {
		term := 4.0 * math.Pow(-1, float64(i)) / float64(2*i+1)
		pi += term
	}
	return pi
}
