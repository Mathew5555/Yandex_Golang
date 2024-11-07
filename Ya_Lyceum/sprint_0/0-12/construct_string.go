package main

import "fmt"

func ConcatStringsAndInt(str1, str2 string, num int) string {
	numS := fmt.Sprintf("%d", num)
	return str1 + str2 + numS
}

func main() {
	fmt.Println(ConcatStringsAndInt("a", "B", 1))
}
