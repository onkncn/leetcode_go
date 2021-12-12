package main

import (
	"fmt"
)

var flag map[byte]bool

func getformat(c byte) byte {
	if c >= '0' && c <= '9' {
		return '0'
	}
	if c == '+' || c == '-' {
		return '+'
	}
	if c == 'e' || c == 'E' {
		return 'e'
	}
	if c == '.' {
		return c
	}
	return 'a'
}

func isNumber(s string) bool {
	flag = make(map[byte]bool)
	if (s[0] != '-' && s[0] != '+' && s[0] != '.' && (s[0] < '0' || s[0] > '9')) || (len(s) == 1 && s[0] == '.') {
		return false
	}
	if s[0] == '.' && (flag['.'] || (flag['e']) || (!flag['0'] && !(1 < len(s) && getformat(s[1]) == '0'))) {
		return false
	}
	flag[getformat(s[0])] = true
	for i := 1; i < len(s); i++ {
		if s[i] == '.' && (flag['.'] || (flag['e']) || (!flag['0'] && !(i+1 < len(s) && getformat(s[i+1]) == '0'))) {
			return false
		}
		if s[i] == 'e' || s[i] == 'E' {
			if flag['e'] || flag['E'] {
				return false
			}
			if !flag['0'] || !((i+1 < len(s) && getformat(s[i+1]) == '0') || (i+2 < len(s) && getformat(s[i+2]) == '0')) {
				return false
			}
		}
		if getformat(s[i]) == '+' && s[i-1] != 'e' {
			return false
		}
		if s[i] != '-' && s[i] != '+' && s[i] != '.' && s[i] != 'e' && s[i] != 'E' && (s[i] < '0' || s[i] > '9') {
			return false
		}
		flag[getformat(s[i])] = true
	}
	return (s[len(s)-1] >= '0' && s[len(s)-1] <= '9') || s[len(s)-1] == '.'
}

func main() {
	ssf := []string{"-e1", ".e1", ".", "1e", "1a", "95a54e53", "e3", "99e2.5", "--6", "-+3", "abc", "+", "-", "a", "e"}
	sst := []string{"2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789", "2", "0089", "-0.1", "+3.14", "4.", "-.9"}
	for _, s := range sst {
		fmt.Println(isNumber(s))
	}
	fmt.Println("fffffffff")
	for _, s := range ssf {
		fmt.Println(isNumber(s))
	}
}
