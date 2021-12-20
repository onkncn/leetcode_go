package main

func sort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	h, t, m := 1, l-1, nums[0]
	for h <= t {
		if nums[h] < m {
			nums[h], nums[h-1] = nums[h-1], nums[h]
			h++
		} else {
			nums[h], nums[t] = nums[t], nums[h]
			t--
		}
	}
	sort(nums[:t])
	sort(nums[h:])
	return nums
}
func mymax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maximumProduct(nums []int) int {
	l := len(nums)
	if l == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	sort(nums)
	return mymax(nums[0]*nums[1]*nums[l-1], nums[l-3]*nums[l-2]*nums[l-1])
}

func main() {

}
