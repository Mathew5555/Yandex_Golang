package main

import (
	"fmt"
)

func main() {
	var num int
	a, b := 0, 1
	var flag bool
	var cnt int = 0
	fmt.Scanln(&num)
	for {
		if !flag && a+b >= num {
			flag = true
		}
		b = a + b
		a = b - a
		if flag {
			fmt.Println(b)
			cnt++
		}
		if cnt == 10 {
			break
		}
	}
}
