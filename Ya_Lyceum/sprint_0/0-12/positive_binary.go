package main

import (
	"errors"
	"fmt"
)

var (
	NegativeNumber = errors.New("negative numbers are not allowed")
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", NegativeNumber
	}
	var res string
	for num > 0 {
		ost := fmt.Sprintf("%d", num%2)
		res = ost + res
		num /= 2
	}
	return res, nil
}
