package main

import (
	"fmt"
)

// Custom error type for negative square root inputs
type ErrNegativeSqrt float64

// Implement the error interface for ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt function using Newton's method
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // Correctly returning error type
	}

	z := float64(1) // Initial guess
	fmt.Printf("Sqrt of %v:\n", x)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}

	return z, nil // Returning the computed value and no error
}

func main() {
	// Corrected calls to Sqrt function
	result, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err) // Properly handles the error message
	} else {
		fmt.Println("Result:", result)
	}
}
