package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z:= float64(1)
	fmt.Printf("Sqrt of %v:\n", x)
    for i := 1; i <= 10; i++ {
        z -= (z*z - x) / (2*z)
        fmt.Printf("iteration %v, value = %v\n", i, z)
    }
    return z 
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
