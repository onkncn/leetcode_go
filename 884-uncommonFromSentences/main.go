package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) (res []string) {
	ss1 := strings.Split(s1, " ")
	ss2 := strings.Split(s2, " ")
	mp := map[string]int{}
	for _, c := range ss1 {
		mp[c]++
	}
	for _, c := range ss2 {
		mp[c]++
	}
	for k, v := range mp {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
