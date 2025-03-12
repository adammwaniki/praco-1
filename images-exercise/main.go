package main

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

/*
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Declaring a slice of slices called `pic` with a length of `dy`.
	// Each inner slice will store values of type uint8.
	pic := make([][]uint8, dy)

	// At the moment the outer slice is empty and each inner slice is nil
	// Looping through the values in the range of the outer slice by index to initialize each inner slice
	for i := range pic {
		// Allocating a slice of uint8 with a length of `dx` at each `pic[i]`.
		pic[i] = make([]uint8, dx)

		// Looping through the inner slice to assign values.
		for j := range pic[i] {
			// Setting the value of each element in both the outer and inner slices.
			pic[i][j] = uint8((i^j))
		}
	}
	// returning the 2D slice with populated values
	return pic
}

func main() {
	pic.Show(Pic)
}
*/