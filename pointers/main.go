package main

import "fmt"

func main() {
	i, j := 42, 2701
	
	var p *int // declares that p is a pointer
	fmt.Printf("Type: %T, zero value: %v\n", p, p) // read the type and value of p
	p = &i         // point to i // changing the assignment to =
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
