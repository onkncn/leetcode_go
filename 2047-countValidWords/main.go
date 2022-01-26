package main

import "strings"

func countValidWords(sentence string) (nums int) {
	s := strings.Split(sentence, " ")
	for _, word := range s {
		times := 0
		l := len(word)
		if l == 0 {
			continue
		}
		flag := true
		for i := 0; i < l; i++ {
			if word[i] == '-' {
				if i == 0 || i == l-1 || times != 0 {
					flag = false
					break
				}
				if !(word[i-1] <= 'z' && word[i-1] >= 'a' && word[i+1] <= 'z' && word[i+1] >= 'a') {
					flag = false
					break
				}
				times++
			} else if word[i] == '!' || word[i] == '.' || word[i] == ',' {
				if i != l-1 {
					flag = false
					break
				}
			} else if word[i] <= 'z' && word[i] >= 'a' {
				continue
			} else {
				flag = false
				break
			}
		}
		if flag {
			nums++
		}
	}
	return nums
}
