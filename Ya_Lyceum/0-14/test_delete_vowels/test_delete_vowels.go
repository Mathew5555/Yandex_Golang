package test_delete_vowels

import "testing"

func DeleteVowels(s string) string {
	return ""
}

type Test struct {
	in       string
	expected string
}

var tests = []Test{
	{"qwerty", "qwrt"},
	{"aao", ""},
	{"tmp", "tmp"},
	{"why", "why"},
}

func TestDeleteVowels(t *testing.T) {
	for _, test := range tests {
		res := DeleteVowels(test.in)
		if res != test.expected {
			t.Errorf("Delete_Vowels(%s)=%s; expected %s", test.in, res, test.expected)
		}
	}
}
