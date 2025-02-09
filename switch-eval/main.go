package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Jumanne ni lini?")
	today := time.Now().Weekday()
	switch time.Tuesday {
	case today + 0:
		fmt.Println("Leo.")
	case today + 1:
		fmt.Println("Kesho.")
	case today + 2:
		fmt.Println("Kesho kutwa.")
	default:
		fmt.Println("Itawahi fika jamani!?")
	}
}
