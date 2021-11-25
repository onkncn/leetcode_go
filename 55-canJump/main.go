package main

import "fmt"

func canJump(nums []int) bool {
	len := len(nums)
	start := 0
	end := start + nums[start]
	nend := end
	for true {
		if end >= len-1 {
			return true
		}
		for i := start + 1; i <= end; i++ {
			if nums[i]+i > nend {
				nend = nums[i] + i
			}
		}
		if nend == end {
			return false
		}
		start = end
		end = nend
	}
	return true
}

func main() {
	nums := []int{1, 2, 3, 0, 0, 0, 0}
	fmt.Println(canJump(nums))
}
