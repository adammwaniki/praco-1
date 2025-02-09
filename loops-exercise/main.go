package main

import "fmt"

// Sqrt calculates the square root of a number using Newton's method.
func sqrt(x float64) float64{
	z:= float64(1) // Initial guess
	fmt.Printf("Sqrt of %v\n", x)
	for i :=1; i<=10; i++{ // 10 iterations
		z -= (z*z - x) / (2*z) // Newton's method
		fmt.Printf("Iteration %v: %v\n", i, z)
	}
	return z
}

func main(){
	fmt.Println(sqrt(0.125))
}