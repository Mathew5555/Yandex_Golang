package main

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{[]int{1, 2}, 2},
	{[]int{-4, 5}, -20},
	{[]int{0, 2}, 0},
}

func TestMultiply(t *testing.T) {
	for i, test := range tests {
		res := Multiply(test.in[0], test.in[1])
		if res != test.out {
			t.Errorf("#%d: Multiply(%v)=%d; want %d", i, test.in, res, test.out)
		}
	}
}
