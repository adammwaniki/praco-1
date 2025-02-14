package main

import "fmt"

// Creating the struct ?
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} // assigning values to the struct
	fmt.Println(v.X, v.Y) // read the values
	v.X = 4 // assign a new value to a particular struct field
	fmt.Println(v.X, v.Y) // read the new values
}
