package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)
	for i := range num {
		if (i+1)%3 == 0 {
			fmt.Println(i + 1)
		}
	}
}
