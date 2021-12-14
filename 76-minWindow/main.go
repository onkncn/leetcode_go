package main

import "fmt"

func minWindow(s string, tt string) string {
	ls, lt := len(s), len(tt)
	if ls < lt {
		return ""
	}
	h, t := 0, 0
	minh, mint := -1, ls
	flag := make(map[byte]int)
	for i, _ := range tt {
		flag[tt[i]]++
	}
	has_nums := make(map[byte]int)
	has_all := func() bool {
		for i, _ := range flag {
			if has_nums[i] < flag[i] {
				return false
			}
		}
		return true
	}
	for t < ls && h < ls {
		for t < ls && !has_all() {
			if flag[s[t]] != 0 {
				has_nums[s[t]]++
			}
			t++
		}
		for h < ls && !has_all() {
			if flag[s[h]] != 0 {
				has_nums[s[h]]--
			}
			h++
		}
		for has_all() {
			if t-h < mint-minh {
				minh, mint = h, t
			}
			if flag[s[h]] != 0 {
				has_nums[s[h]]--
			}
			h++
		}
		if h == ls {
			break
		}
		if flag[s[h]] != 0 {
			has_nums[s[h]]--
		}
		h++
	}
	if minh == -1 {
		return ""
	}
	return s[minh:mint]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
