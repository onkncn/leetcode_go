package main

import (
	"fmt"
	"sort"
)

func search(nums []int, target int) bool {
	nlen := len(nums)
	if nlen < 10 {
		for i := 0; i < nlen; i++ {
			if nums[i] == target {
				return true
			}
			return false
		}
	}
	if target == nums[0] {
		return true
	}
	index := sort.Search(nlen, func(index int) bool {
		return nums[index-1] <= nums[index] && nums[index+1] > nums[index]
	})
	nums1 := nums[0:index]
	i := sort.Search(nlen, func(index int) bool {
		return nums1[index] >= target
	})
	if nums1[i] == target {
		return true
	}
	nums2 := nums[index:nlen]
	i = sort.Search(nlen, func(index int) bool {
		return nums2[index] >= target
	})
	return nums2[i] == target
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
