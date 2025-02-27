package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Creating a slice of words from the `s` using whitespace as the delimiter.
	words := strings.Fields(s)
	
	// Creating an empty map to store word frequencies.
	// The keys represent unique words, and the values store their respective 		// counts.
	wordMap := make(map[string]int)
	
	// Looping through the values in the slice of words
	for _, v := range words{
		// Incrementing the count for the word in the map.
		wordMap[v] = wordMap[v] + 1
		// alternatively use: wordMap[v]++
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}


// Below are the expected outputs from the test file for the above code:
/*

PASS
 f("I am learning Go!") = 
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") = 
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") = 
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") = 
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}

*/