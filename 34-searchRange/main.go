package main

func binarySearch_p(nums []int, target int) int {
	h, t := 0, len(nums)
	m := (h + t) / 2
	for h < t {
		if nums[m] == target && (m < 1 || nums[m-1] != target) {
			return m
		}
		if nums[m] >= target {
			t, m = m, (h+m)/2
		} else {
			h, m = m+1, (t+m)/2
		}
	}
	return h
}
func binarySearch_q(nums []int, target int) int {
	h, t := 0, len(nums)
	m := (h + t) / 2
	for h < t {
		if nums[m] == target && (m > len(nums)-2 || nums[m+1] != target) {
			return m
		}
		if nums[m] > target {
			t, m = m, (h+m)/2
		} else {
			h, m = m+1, (t+m)/2
		}
	}
	return h
}
func searchRange(nums []int, target int) []int {
	res := []int{binarySearch_p(nums, target), binarySearch_q(nums, target)}
	if len(nums) == 0 || res[0] >= len(nums) || nums[res[0]] != target {
		return []int{-1, -1}
	}
	return res
}
func main() {

}
