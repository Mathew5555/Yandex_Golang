package main

import (
	"errors"
	"fmt"
)

var (
	PositionOutOfRange = errors.New("position out of range")
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	runeStr := []rune(str)
	if position >= len(runeStr) {
		return 0, PositionOutOfRange
	}
	return runeStr[position], nil
}

func main() {
	fmt.Println(rune("О"[0]))
	fmt.Println(GetCharacterAtPosition("Оле ола вперёд Спартак Москва", 0))
}
