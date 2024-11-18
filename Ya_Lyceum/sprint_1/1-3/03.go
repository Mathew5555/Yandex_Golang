package main

import (
	"fmt"
	"strings"
)

type UpperWriter struct {
	UpperString string
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func (w *UpperWriter) Write(p []byte) (n int, err error) {
	w.UpperString = strings.ToUpper(string(p))
	return len(p), err
}

func main() {
	// Create a new UpperWriter.
	w := UpperWriter{}

	// Write some data to the UpperWriter.
	if _, err := w.Write([]byte("hello world")); err != nil {
		fmt.Println(err)
	}

	// Print the UpperString field.
	fmt.Println(w.UpperString) // Output: "HELLO WORLD"
}
