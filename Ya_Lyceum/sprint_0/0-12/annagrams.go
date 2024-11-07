package main

import (
	"fmt"
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	lowerStr1 := strings.ToLower(str1)
	lowerStr2 := strings.ToLower(str2)
	splitStr1 := strings.Split(lowerStr1, "")
	splitStr2 := strings.Split(lowerStr2, "")
	sort.Strings(splitStr1)
	sort.Strings(splitStr2)
	for i := 0; i < len(splitStr1); i++ {
		if splitStr1[i] != splitStr2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(AreAnagrams("Кабан", "банка"))
}
