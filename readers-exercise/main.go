package main

import (
	"golang.org/x/tour/reader"
)

// Define MyReader type
type MyReader struct{}

// Implement the Read method for MyReader
func (mr MyReader) Read(b []byte) (int, error) {
	// Fill the provided slice with 'A'
	for i := range b {
		b[i] = 'A'
	}
	// Return the number of bytes filled and no error
	return len(b), nil
}

func main() {
	// Validate that MyReader implements io.Reader correctly
	reader.Validate(MyReader{})
}
