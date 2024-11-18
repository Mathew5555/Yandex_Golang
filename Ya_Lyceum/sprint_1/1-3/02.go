package main

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	p := make([]byte, 1024)
	bytesRead, err := r.Read(p)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(p[:bytesRead]), nil
}
