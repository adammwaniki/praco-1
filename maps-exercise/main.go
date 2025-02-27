package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Creating a slice of words from the string s by splitting at whitespace
	words := strings.Fields(s)
	count := 0
	for i, v := range words{
		
	}
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
