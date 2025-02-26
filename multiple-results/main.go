package main

import "fmt"

// Creating a function that returns two strings whose order has been swapped.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}