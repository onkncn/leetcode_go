package main

import "strconv"

func optimalDivision(nums []int) string {
	res := ""
	res += strconv.Itoa(nums[0])
	l := len(nums)
	if l == 1 {
		return res
	}
	if l == 2 {
		return res + "/" + strconv.Itoa(nums[1])
	}
	res += "/("
	i := 1
	for i = 1; i < l-1; i++ {
		res += strconv.Itoa(nums[i]) + "/"
	}
	return res + strconv.Itoa(nums[i]) + ")"
}

func main() {
	nums := []int{2, 3, 4}
	optimalDivision(nums)
}
