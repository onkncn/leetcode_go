package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	nlen := len(nums)
	res := [][]int{{}}
	sort.Ints(nums)
	for i := 0; i < nlen; i++ {
		times := 1
		for ; i+1 < nlen && nums[i+1] == nums[i]; i++ {
			times++
		}
		lr := len(res)
		tres := make([][]int, lr*times)
		for j := 0; j < len(res); j++ {
			tres[j] = make([]int, 1+len(res[j]))
			copy(tres[j], res[j])
			tres[j][len(res[j])] = nums[i]
		}
		for k := 1; k < times; k++ {
			for j := 0; j < len(res); j++ {
				tres[j+k*lr] = make([]int, len(res[j])+k+1)
				copy(tres[j+k*lr], tres[j+k*lr-lr])
				tres[j+k*lr][len(res[j])+k] = nums[i]
			}
		}
		res = append(res, tres...)
	}
	return res
}

func main() {
	fmt.Println(subsetsWithDup([]int{3, 2, 2}))
}
