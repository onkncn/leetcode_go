package main

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 1 {
		return 1
	}
	s += "."
	flag := make(map[byte]bool)
	h, t := 0, 0
	max := 0
	for t < l {
		if !flag[s[t]] {
			flag[s[t]] = true
			t++
		} else {
			flag[s[h]] = false
			h++
		}
		if t-h > max {
			max = t - h
		}
	}
	return max
}
