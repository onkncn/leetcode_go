package main

func findPeakElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[l-1] > nums[l-2] {
		return l - 1
	}
	l, r := 1, l-1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] && nums[m] > nums[m-1] {
			return m
		}
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return 0
}
