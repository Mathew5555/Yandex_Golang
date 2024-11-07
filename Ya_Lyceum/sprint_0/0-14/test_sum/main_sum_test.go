package test_sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "positive values",
			values:   []int{1, 2},
			expected: 3,
		},
		{
			name:     "negative values",
			values:   []int{-1, -2},
			expected: -3,
		},
		{
			name:     "mixed values",
			values:   []int{-1, 2},
			expected: 1,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.values[0], tc.values[1])
			if result != tc.expected {
				t.Errorf("Sum(%v) = %v; expected %v", tc.values, result, tc.expected)
			}
		})
	}
}
