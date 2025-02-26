/*
Implement Pic. It should return a slice of length dy, 
each element of which is a slice of dx 8-bit unsigned integers. 
When you run the program, it will display your picture, 
interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. 
Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

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


// Alternative Solution
/*
func Pic(dx, dy int) [][]uint8 {
	// Declaring a slice of slices
	ss := make([][]uint8, dy)
	    // Looping through the outer slice
		for y := 0; y < dy; y++ {
			s := make([]uint8, dx)
			for x := 0; x < dx; x++ {
				s[x] = uint8(x*y)
			}
			ss[y] = s
		}
		return ss
}

func main() {
	pic.Show(Pic)
}
*/