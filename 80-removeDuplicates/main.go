package main

func removeDuplicates(nums []int) int {
	nlen := len(nums)
	if nlen <= 2 {
		return nlen
	}
	index := 1
	times := 1
	for i := 1; i < nlen; i++ {
		if nums[index-1] == nums[i] {
			times++
		} else {
			times = 1
		}
		if times <= 2 {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
func main() {

}
