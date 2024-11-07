package main

import (
	"errors"
	"fmt"
)

var (
	ErrorDivisionByZero = errors.New("division by zero is not allowed")
)

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, ErrorDivisionByZero
	}
	return float64(a / b), nil
}

func main() {
	fmt.Println(DivideIntegers(5, 1))
}
