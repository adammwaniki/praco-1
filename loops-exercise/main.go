package main

import (
	"fmt"
	"time"
)

// Sqrt calculates the square root of a number using Newton's method.
func sqrt(x float64) float64{
	startTime := time.Now() // Start time measurement
	var z float64 = 1 // Initial guess
	fmt.Printf("Sqrt of %v\n", x)
	for i :=1; i<=10; i++{ // 10 iterations
		z -= (z*z - x) / (2*z) // Newton's method
		fmt.Printf("Iteration %v: %v\n", i, z)
	}
	elapsedTime := time.Since(startTime) // End time measurement and calculate elapsed time
	fmt.Printf("Execution time: %v\n", elapsedTime)
	return z
}

func main(){
	fmt.Println(sqrt(0.125), sqrt(2), sqrt(4), sqrt(16))
}