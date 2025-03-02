package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// Initialize the first two numbers in the Fibonacci sequence
	n2, n1 := 0, 1
	return func() int {
		result := n2 // Store the current Fibonacci number including the initial 0
		// Update the sequence: 
		// n2 takes the value of n1, and n1 becomes the sum of both
		n2, n1 = n1, n1+n2
		return result // Return the Fibonacci number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
