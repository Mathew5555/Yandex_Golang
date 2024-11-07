package main

func Join(nums1, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))
	for i := 0; i < len(nums1); i++ {
		res[i] = nums1[i]
	}
	for i := 0; i < len(nums2); i++ {
		res[len(nums1)+i] = nums2[i]
	}
	return res
}

func main() {
}
