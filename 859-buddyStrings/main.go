package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	has := [26]int{}
	flag := 0
	flag2 := false
	swapi, swapj := -1, -1
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			flag++
			if flag > 2 {
				return false
			}
			if swapi == -1 {
				swapi = i
			} else if swapj == -1 {
				swapj = i
			}
		} else {
			if flag2 == false {
				has[s[i]-'a']++
				if has[s[i]-'a'] >= 2 {
					flag2 = true
				}
			}
		}
	}
	if flag == 1 {
		return false
	}
	if flag == 2 {
		return s[swapi] == goal[swapj] && s[swapj] == goal[swapi]
	}
	return flag2
}

func main() {
	fmt.Println(buddyStrings("ab", "babbb"))
}
