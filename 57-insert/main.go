package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func insert(intervals [][]int, newInterval []int) [][]int {
	leni := len(intervals)
	if leni == 0 {
		return [][]int{newInterval}
	}
	res := [][]int{}
	i := 0
	j := 0
	for i < leni && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	if i == leni || intervals[i][0] > newInterval[1] {
		res = append(res, newInterval)
	} else {
		res = append(res, []int{min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1])})
	}
	j = i
	for j < leni && res[i][1] >= intervals[j][0] {
		res[i][1] = max(res[i][1], intervals[j][1])
		j++
	}
	i = j
	for ; i < leni; i++ {
		res = append(res, intervals[i])
	}
	return res
}

/*
x1 10
5 y2

*/

func main() {
	iii := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	jjj := []int{4, 8}
	//fmt.Println(insert(iii,jjj))
	iii = [][]int{{1, 5}}
	jjj = []int{1, 3}
	//fmt.Println(insert(iii,jjj))
	iii = [][]int{{1, 5}}
	jjj = []int{0, 0}
	fmt.Println(insert(iii, jjj))
}
