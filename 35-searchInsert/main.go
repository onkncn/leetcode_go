package main

func searchInsert(nums []int, target int) int {
	h, t := 0, len(nums)
	m := (h + t) / 2
	for h < t {
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			m, t = (h+m)/2, m
		} else {
			h, m = m+1, (m+t)/2
		}
	}
	return h
}
func main() {

}
