package main

import "fmt"

func Clean(nums []int, x int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}
	return nums
}

func main() {
	a := []int{5, 5}
	fmt.Println(Clean(a, 5))
}
