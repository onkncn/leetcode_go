package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	la, lb := len(a), len(b)
	res := a
	lr := len(res)
	for lr < lb {
		res += a
		lr += la
	}
	if strings.Index(res, b) != -1 {
		return lr / la
	}
	res += a
	lr += la
	if strings.Index(res, b) != -1 {
		return lr / la
	}
	return -1
}

func main() {
	fmt.Println(repeatedStringMatch("abab", "aba"))
}
