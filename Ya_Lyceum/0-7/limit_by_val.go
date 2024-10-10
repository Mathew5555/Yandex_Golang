package main

import (
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("n неотрицательное")
	} else if nums == nil {
		return nil, fmt.Errorf("массив непустой")
	}
	res := make([]int, 0, n)
	for _, el := range nums {
		if el < limit {
			res = append(res, el)
			if len(res) == n {
				break
			}
		}
	}
	return res, nil
}

func TestUnderLimit() {
	type test struct {
		nums      []int
		n         int
		limit     int
		expected  []int
		wantError bool
	}

	tests := []test{
		{
			nums:      []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12},
			n:         5,
			limit:     3,
			expected:  []int{2, 0, 2},
			wantError: false,
		},
		{
			nums:      nil,
			wantError: true,
		},
		{
			nums:      []int{},
			n:         5,
			limit:     3,
			expected:  []int{},
			wantError: false,
		},
		{
			nums:      []int{3, 5, 6},
			n:         5,
			limit:     10,
			expected:  []int{3, 5, 6},
			wantError: false,
		},
		{
			nums:      []int{-13, 0, 6},
			n:         1,
			limit:     -5,
			expected:  []int{-13},
			wantError: false,
		},
		{
			nums:      []int{},
			n:         -1,
			limit:     5,
			wantError: true,
		},
	}

	for _, tc := range tests {
		_, err := UnderLimit(tc.nums, tc.limit, tc.n)
		fmt.Println(tc.nums)
		fmt.Println(err)
		if tc.wantError {
			if err == nil {
				fmt.Println("expec... File is too long to be displayed fully")
			}
		}
	}
}

func main() {
	TestUnderLimit()
}
