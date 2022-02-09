package main

import "fmt"

func maxProduct(nums []int) int {
	l := len(nums)
	maxnums := make([]int, l)
	minnums := make([]int, l)
	if l == 1 {
		return nums[0]
	}
	maxnums[0] = nums[0]
	minnums[0] = nums[0]
	max := nums[0]
	for i := 1; i < l; i++ {
		maxnums[i], minnums[i] = compare3(nums[i], minnums[i-1]*nums[i], maxnums[i-1]*nums[i])
		if maxnums[i] > max {
			max = maxnums[i]
		}
		fmt.Println(maxnums[i], minnums[i])
	}
	return max
}
func compare2(a, b int) (big, small int) {
	if a > b {
		return a, b
	}
	return b, a
}
func compare3(a, b, c int) (big, small int) {
	a, b = compare2(a, b)
	big, _ = compare2(a, c)
	_, small = compare2(b, c)
	return big, small
}
func main() {
	fmt.Println(maxProduct([]int{-2, 3, -4}))
}
