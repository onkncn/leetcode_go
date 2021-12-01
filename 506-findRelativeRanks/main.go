package main

import (
	"fmt"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	mp := map[int]int{}
	ls := len(score)
	t := make([]int, ls)
	copy(t, score)
	res := []string{}
	qsort(t)
	fmt.Println(t)
	for i := 0; i < ls; i++ {
		mp[t[i]] = i + 1
	}
	for i := 0; i < ls; i++ {
		if mp[score[i]] == 1 {
			res = append(res, "Gold Medal")
		} else if mp[score[i]] == 2 {
			res = append(res, "Silver Medal")
		} else if mp[score[i]] == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(mp[score[i]]))
		}
	}
	return res
}

func qsort(a []int) []int {
	l := len(a)
	if l < 2 {
		return a
	}
	h, t, m := 1, l-1, a[0]
	for h <= t {
		if a[h] > m {
			a[h], a[h-1] = a[h-1], a[h]
			h++
		} else {
			a[h], a[t] = a[t], a[h]
			t--
		}
	}
	qsort(a[:t])
	qsort(a[h:])
	return a
}

func main() {

}
