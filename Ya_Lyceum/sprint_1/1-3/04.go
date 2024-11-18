package main

import "io"

func Copy(r io.Reader, w io.Writer, n uint) error {
	p := make([]byte, n)
	bytesRead, err := r.Read(p)
	if err != nil && err != io.EOF {
		return err
	}
	_, err = w.Write(p[:bytesRead])
	return err
}
