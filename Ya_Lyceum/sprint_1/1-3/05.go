package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	EmptyString = errors.New("Empty string")
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return false, EmptyString
	}
	p := make([]byte, 1024)
	n, err := r.Read(p)
	if err != nil && err != io.EOF {
		return false, nil
	}
	text := string(p[:n])
	seqText := string(seq)
	index := strings.Index(text, seqText)
	return index != -1, nil
}

func main() {
	r := bytes.NewReader([]byte("Hello, world!"))
	seq := []byte("")
	found, err := Contains(r, seq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(found) // Output: true
}
