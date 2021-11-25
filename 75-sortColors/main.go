package main

import "fmt"

func sortColors(nums []int) {
	nlen := len(nums)
	index0, index2 := 0, nlen-1
	for i := 0; i < nlen; i++ {
		if i > index2 {
			break
		}
		if nums[i] == 0 && i != index0 {
			nums[i], nums[index0] = nums[index0], nums[i]
			index0++
			i--
		} else if nums[i] == 2 {
			nums[i], nums[index2] = nums[index2], nums[i]
			index2--
			i--
		}
	}
}
func main() {
	nums := []int{2, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
