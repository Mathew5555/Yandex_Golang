package main

import (
	"errors"
	"testing"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

type Test struct {
	in            []byte
	expected      int
	expectedError error
}

var tests = []Test{
	{
		[]byte("Hello, 世界"),
		9,
		nil,
	},
	{
		[]byte{0xff, 0x1a},
		0,
		ErrInvalidUTF8,
	},
	{
		[]byte("Hi"),
		2,
		nil,
	},
}

func TestGetUTFLength(t *testing.T) {
	for _, test := range tests {
		res, err := GetUTFLength(test.in)
		if res != test.expected && err != test.expectedError {
			if test.expectedError != nil {
				t.Errorf("GetUTFLength(%v)=%s; expected %s", test.in, err, test.expectedError)
			} else {
				t.Errorf("GetUTFLength(%v)=%d; expected %v", test.in, res, test.expected)
			}
		}
	}
}
