package main

import "fmt"

func IsPowerOfTwoRecursive(N int) {
	if N%2 == 0 {
		IsPowerOfTwoRecursive(N / 2)
	} else if N == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
}
