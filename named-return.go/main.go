package main

import "fmt"

func spit(z int) (x, y int) {
	x = z * 4/9
	y = z - x
	return
}

func main() {
	fmt.Println(spit(17))
}