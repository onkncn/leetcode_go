package main

import "fmt"

var flag bool = false

func isInterleave(s1 string, s2 string, s3 string) bool {
	flag = false
	return help(s1, s2, s3)
}
func help(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 || flag {
		return false
	}
	if l1+l2 == 1 && s1+s2 != s3 {
		flag = true
	}
	if l1 == 0 {
		return s2 == s3
	}
	if l2 == 0 {
		return s1 == s3
	}
	if s1[0] == s3[0] {
		if help(s1[1:l1], s2, s3[1:l3]) {
			return true
		}
	}
	if s2[0] == s3[0] {
		if help(s1, s2[1:l2], s3[1:l3]) {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(isInterleave("a", "b", "ab"))
}
