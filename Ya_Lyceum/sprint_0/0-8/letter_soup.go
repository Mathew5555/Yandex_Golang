package main

import (
	"fmt"
	"sort"
)

func Permutations(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}

	var lsPermutations []string
	for ind, letter := range input {
		letterStr := string(letter)
		var inpWithoutLetter = input[:ind] + input[ind+1:]
		var supPerm = Permutations(inpWithoutLetter)
		for _, el := range supPerm {
			lsPermutations = append(lsPermutations, letterStr+el)
		}
	}
	sort.Strings(lsPermutations)
	return lsPermutations
}

func main() {
	str := "cab"
	permutations := Permutations(str)
	fmt.Println("Permutations:", permutations)
}
