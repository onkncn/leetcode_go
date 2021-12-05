package main

import "fmt"

func bSt(nums []int, target int) int {
	h, t := 0, len(nums)
	s, e := nums[h], nums[t-1]
	if s < e {
		return 0
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
func search(nums []int, target int) int {
	t := bSt(nums, target)
	fmt.Println(t)
	l := len(nums)
	a := bS(nums, 0, t, target)
	if a < l && nums[a] == target {
		return a
	}
	a = bS(nums, t, len(nums), target)
	if a < l && nums[a] == target {
		return a
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 9))
}
