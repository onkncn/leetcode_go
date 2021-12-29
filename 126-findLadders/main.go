package main

import "fmt"

var l int

func just_1_difffrent(a, b string) bool {
	times := 0
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			times++
		}
	}
	return times == 1
}

var before map[string][]string

func findLadders(beginWord string, endWord string, wordList []string) (res [][]string) {
	mp := make(map[string]int)
	queue := []string{beginWord}
	mp[beginWord] = 0
	lq, times := 1, 1
	l = len(beginWord)
	before = make(map[string][]string)
	for lq != 0 {
		times++
		final := false
		for i := 0; i < lq; i++ {
			for j := 0; j < len(wordList); j++ {
				if just_1_difffrent(queue[i], wordList[j]) {
					if mp[wordList[j]] == 0 {
						before[wordList[j]] = append(before[wordList[j]], queue[i])
						if wordList[j] == endWord {
							final = true
						}
						queue = append(queue, wordList[j])
						mp[wordList[j]] = times
					} else if mp[wordList[j]] == times {
						before[wordList[j]] = append(before[wordList[j]], queue[i])
					}
				}
			}
		}
		if final {
			break
		}
		queue = queue[lq:]
		lq = len(queue)
	}
	has := []string{}
	var dfs func(now string, has []string)
	dfs = func(now string, has []string) {
		has = append([]string{now}, has...)
		if now == beginWord {
			res = append(res, has)
		} else {
			for _, s := range before[now] {
				dfs(s, has)
			}
		}
	}
	dfs(endWord, has)
	return
}

func main() {
	fmt.Println(findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
}
