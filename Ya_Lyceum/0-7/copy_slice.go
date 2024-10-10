package main

import "fmt"

func SliceCopy(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = nums[i]
	}
	//res = append(res, nums...)
	fmt.Println(cap(res), len(res))
	return res
}

func main() {
	a := []int{1, 2, 3}
	SliceCopy(a)
}
