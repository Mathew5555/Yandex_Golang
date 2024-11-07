package main

import (
	"errors"
	"fmt"
)

var (
	NegativeFactorial = errors.New("factorial is not defined for negative numbers")
)

func Factorial(n int) (int, error) {
	switch {
	case n < 0:
		return 0, NegativeFactorial
	case n == 0:
		return 1, nil
	default:
		prev, _ := Factorial(n - 1)
		return n * prev, nil
	}
}

func main() {
	fmt.Println(Factorial(-1))
}
