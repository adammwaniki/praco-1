package main

import (
	"fmt"
	"time"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	start := time.Now() // Store start time

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // (*p).X = 1e9

	fmt.Println(v)

	elapsed := time.Since(start) // Compute elapsed time
	fmt.Println(elapsed)
}
