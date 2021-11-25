package main

import (
	"fmt"
)

var inters [][]int

func merge2(a, b []int) [][]int {
	if a[0] <= b[0] {
		if a[1] < b[0] {
			return [][]int{a, b}
		} else if a[1] <= b[1] {
			return [][]int{{a[0], b[1]}}
		} else {
			return [][]int{a}
		}
	} else {
		return merge2(b, a)
	}
}
func compare(a, b []int) bool {
	if a[0] == b[0] {
		return a[1] < b[1]
	}
	return a[0] > b[0]
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	leni := len(intervals)
	for i := 0; i < leni; i++ {
		for j := i + 1; j < leni; j++ {
			if compare(intervals[i], intervals[j]) {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	for i := 0; i < leni; i++ {
		t := intervals[i]
		j := i + 1
		for ; j < leni; j++ {
			ti := merge2(t, intervals[j])
			if len(ti) == 2 {
				i = j - 1
				break
			}
			t = ti[0]
		}
		i = j - 1
		res = append(res, t)
	}
	return res
}
func main() {
	inters = [][]int{{1, 9}, {3, 4}, {1, 8}, {3, 5}}
	fmt.Println(merge(inters))
}
