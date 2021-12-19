package main

func minSubArrayLen(target int, nums []int) int {
	res := 111111
	h, t, l, sum := 0, 0, len(nums), 0
	nums = append(nums, -111111)
	for t <= l {
		if sum < target {
			sum += nums[t]
			t++
		} else {
			if res > t-h {
				res = t - h
			}
			sum -= nums[h]
			h++
		}
	}
	if res == 111111 {
		return 0
	}
	return res
}

func main() {

}
