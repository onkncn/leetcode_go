package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func sort(a []int) []int {
	l := len(a)
	if l < 2 {
		return a
	}
	h, t, m := 1, l-1, a[0]
	for h <= t {
		if a[h] < m {
			a[h], a[h-1] = a[h-1], a[h]
			h++
		} else {
			a[h], a[t] = a[t], a[h]
			t--
		}
	}
	sort(a[:t])
	sort(a[h:])
	return a
}
func largestSumAfterKNegations(nums []int, k int) int {
	amin, ln, sum := 111, len(nums), 0
	nums = sort(nums)
	fmt.Println(nums)
	for i := 0; i < ln; i++ {
		if nums[i] < 0 && k > 0 {
			sum += -nums[i]
			amin = min(amin, -nums[i])
			k--
		} else {
			sum += nums[i]
			amin = min(amin, nums[i])
		}
	}
	fmt.Println(sum)
	if k%2 == 1 {
		sum = sum - 2*amin
	}
	return sum
}
func main() {

}
