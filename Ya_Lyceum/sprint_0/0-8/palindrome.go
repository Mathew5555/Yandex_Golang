package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {
	input = strings.Replace(input, " ", "", -1)
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("ac a"))
}
