package main

import (
	"fmt"
	"unicode"
)

func isLatin(input string) bool {
	for _, el := range input {
		if !unicode.In(el, unicode.Latin) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isLatin("jjl"))
}
