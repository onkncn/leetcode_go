package main

import "fmt"

func isReverse(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
func partition(s string) [][]string {
	if s == "" {
		return [][]string{}
	}
	res := [][]string{}
	if isReverse(s) {
		res = append(res, []string{s})
	}
	l := len(s)
	for i := 1; i < len(s); i++ {
		if isReverse(s[0:i]) {
			for _, ts := range partition(s[i:l]) {
				t := []string{s[0:i]}
				t = append(t, ts...)
				res = append(res, t)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(partition("aab"))
}
