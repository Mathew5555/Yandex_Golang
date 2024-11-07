package rpn

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrorDivisionByZero = errors.New("division by zero is not allowed")
	UnexpectedElements  = errors.New("unexpected elements in expression")
	IncorrectExpression = errors.New("incorrect expression")
)

func strip(s string) string {
	var l, r int = 0, len(s)
	for l < r && s[l] == ' ' {
		l++
	}
	for l < r && s[r-1] == ' ' {
		r--
	}
	return s[l:r]
}

func split(s, signs string) (byte, string, string) {
	var balance int = 0
	for i := len(s) - 1; i >= 0; i-- {
		if balance == 0 && strings.Contains(signs, string(s[i])) {
			return s[i], s[:i], s[i+1:]
		} else if s[i] == '(' {
			balance--
		} else if s[i] == ')' {
			balance++
		}
	}
	return 0, "", ""
}

func number(s string) (float64, error) {
	s = strip(s)
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, UnexpectedElements
	}
	return float64(num), nil
}

func factor(s string) (float64, error) {
	var res float64
	var err error
	s = strip(s)
	if s == "" {
		return 0, IncorrectExpression
	}
	if s[0] == '+' {
		res, err = factor(s[1:])
	} else if s[0] == '-' {
		res, err = factor(s[1:])
		res *= -1
	} else if s[0] == '(' {
		res, err = Calc(s[1 : len(s)-1])
	} else {
		res, err = number(s)
	}
	return res, err
}

func item(s string) (float64, error) {
	var op byte
	var l, r string
	op, l, r = split(s, "*/")
	if op != 0 {
		res1, err1 := item(l)
		res2, err2 := factor(r)
		if err1 != nil {
			return 0, err1
		} else if err2 != nil {
			return 0, err2
		}
		if op == '*' {
			return res1 * res2, nil
		}
		if res2 == 0 {
			return 0, ErrorDivisionByZero
		}
		return res1 / res2, nil
	}
	return factor(s)
}

func Calc(s string) (float64, error) {
	var op byte
	var l, r string
	op, l, r = split(s, "+-")
	if op != 0 {
		res1, err1 := Calc(l)
		res2, err2 := item(r)
		if err1 != nil {
			return 0, err1
		} else if err2 != nil {
			return 0, err2
		}
		if op == '+' {
			return res1 + res2, nil
		}
		return res1 - res2, nil
	}
	return item(s)
}

//func main() {
//	var exp string
//	fmt.Scanln(&exp)
//	ans, err := Calc(exp)
//	fmt.Println(ans, err)
//	//var flag bool = false
//}
