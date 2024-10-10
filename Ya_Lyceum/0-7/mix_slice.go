package main

import "fmt"

func Mix(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums)/2; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[len(nums)/2+i]
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(Mix(a))
}
