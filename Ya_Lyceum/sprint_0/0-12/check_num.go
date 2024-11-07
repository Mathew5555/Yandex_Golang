package main

import (
	"errors"
	"strconv"
)

var (
	NotInteger = errors.New("invalid input, please provide two integers")
)

func SumTwoIntegers(a, b string) (int, error) {
	aNum, err := strconv.Atoi(a)
	bNum, err2 := strconv.Atoi(b)
	if err != nil || err2 != nil {
		return 0, NotInteger
	}
	return aNum + bNum, nil
}
