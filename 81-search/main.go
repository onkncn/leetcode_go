package main

import (
	"fmt"
)

func bSt(nums []int, target int) int {
	h, t := 0, len(nums)
	s, e := nums[h], nums[t-1]
	if s == e {
		for h < t && nums[t-1] == s && nums[h] == s {
			h++
			t--
			fmt.Println(h, "   ", t)
		}
		if h >= t {
			return 0
		}
		if nums[h] < nums[t-1] {
			return t
		}
		s, e = nums[h], nums[t-1]
	}
	for h < t {
		m := (h + t) / 2
		if m < 1 || nums[m] < nums[m-1] {
			return m
		} else if nums[m] <= e {
			t = m
		} else {
			h = m + 1
		}
	}
	return h
}
func bS(nums []int, h, t, target int) int {
	for h < t {
		m := (h + t) / 2
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			t = m
		} else {
			h = m + 1
		}
	}
	return h
}
func search(nums []int, target int) bool {
	t := bSt(nums, target)
	l := len(nums)
	a := bS(nums, 0, t, target)
	if a < l && nums[a] == target {
		return true
	}
	a = bS(nums, t, len(nums), target)
	if a < l && nums[a] == target {
		return true
	}
	return false
}
func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target := 13
	fmt.Println(search(nums, target))
}
