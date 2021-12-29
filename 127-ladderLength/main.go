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
func ladderLength(beginWord string, endWord string, wordList []string) int {
	mp := make(map[string]int)
	begin_queue := []string{beginWord}
	end_queue := []string{endWord}
	mp[beginWord] = 1
	mp[endWord] = -1
	blq, elq, times := 1, 1, 1
	l = len(beginWord)
	flag := false
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			flag = true
		}
	}
	if !flag {
		return 0
	}
	if just_1_difffrent(beginWord, endWord) {
		return 2
	}
	for blq != 0 && elq != 0 {
		times++
		if blq < elq {
			for i := 0; i < blq; i++ {
				for j := 0; j < len(wordList); j++ {
					if mp[wordList[j]] != 1 && just_1_difffrent(begin_queue[i], wordList[j]) {
						if mp[wordList[j]] == -1 {
							return times
						}
						begin_queue = append(begin_queue, wordList[j])
						mp[wordList[j]] = 1
					}
				}
			}
			begin_queue = begin_queue[blq:]
			blq = len(begin_queue)
		} else {
			for i := 0; i < elq; i++ {
				for j := 0; j < len(wordList); j++ {
					if mp[wordList[j]] != -1 && just_1_difffrent(end_queue[i], wordList[j]) {
						if mp[wordList[j]] == 1 {
							return times
						}
						end_queue = append(end_queue, wordList[j])
						mp[wordList[j]] = -1
					}
				}
			}
			end_queue = end_queue[elq:]
			elq = len(end_queue)
		}
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("a", "b", []string{"b", "c", "d"}))
}
