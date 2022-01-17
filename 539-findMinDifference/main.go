package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	l := len(timePoints)
	if l >= 1440 {
		return 0
	}
	sort.Strings(timePoints)
	min := 1440
	ph, _ := strconv.Atoi(timePoints[l-1][:2])
	ph -= 24
	pm, _ := strconv.Atoi(timePoints[l-1][3:])
	for i := 0; i < l; i++ {
		h, _ := strconv.Atoi(timePoints[i][:2])
		m, _ := strconv.Atoi(timePoints[i][3:])
		t := (h-ph)*60 + m - pm
		if t < min {
			min = t
		}
		ph, pm = h, m
	}
	return min
}

func main() {
	fmt.Println("123")
}
