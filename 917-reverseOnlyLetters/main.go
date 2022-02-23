package main

import "fmt"

func reverseOnlyLetters(s string) string {
	l, r := 0, len(s)-1
	res := make([]byte, len(s))
	copy(res, s)
	for {
		for l < r && !((res[l] >= 'a' && res[l] <= 'z') || (res[l] >= 'A' && res[l] <= 'Z')) {
			l++
		}
		for l < r && !((res[r] >= 'a' && res[r] <= 'z') || (res[r] >= 'A' && res[r] <= 'Z')) {
			r--
		}
		if l >= r {
			break
		}
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}

func main() {
	fmt.Println(reverseOnlyLetters("sdhgkjahsgkh42123dsf"))
}
