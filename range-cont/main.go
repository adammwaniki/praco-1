package main

import "fmt"

func main() {
	// Create a slice of integers with length 10 and initialized with zero values.
	pow := make([]int, 10)

	// Loop over the slice by index (i).
	for i := range pow {
		// Set the value at index i to 2 raised to the power of i.
		// 1 << uint(i) is equivalent to 2**i.
		pow[i] = 1 << uint(i)
	}

	// Loop through the slice and print each value.
	// The blank identifier "_" is used to ignore the index.
	for _, value := range pow {
		// Print the value in the slice.
		fmt.Printf("%d\n", value)
	}
}
