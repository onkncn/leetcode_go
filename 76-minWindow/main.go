package main

import "fmt"

func minWindow(s string, tt string) string {
	ls, lt := len(s), len(tt)
	if ls < lt {
		return ""
	}
	h, t := 0, 0
	minh, mint := 0, ls
	flag := [222]int{}
	for _, c := range tt {
		flag[c]++
	}
	has_nums := [222]int{}
	for h < ls && t < ls {
		for has_nums != flag {
			if flag[s[t]] != 0 {
				has_nums[s[t]]++
				if t+1-h < mint-minh {
					minh, mint = h, t+1
					break
				}
			}
			t++
		}
		for h < ls && flag[s[h]] == 0 {
			h++
		}
		if h < ls {
			has_nums[s[h]]--
		}
		h++
	}
	if mint != ls {
		return s[minh : mint+1]
	}
	return ""
}

func main() {
	fmt.Println(minWindow("ABCCCCDFS", "AFS"))
}
