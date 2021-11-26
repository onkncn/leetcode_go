package main

import "fmt"

var res [][]int

func dfs(nums []int, t []int) {
	res = append(res, t)
	nlen := len(nums)
	if nlen == 0 {
		return
	}
	for i := 0; i < nlen; i++ {
		addt := make([]int, len(t)+1)
		copy(addt, t)
		addt[len(t)] = nums[i]
		fmt.Println(addt)
		dfs(nums[i+1:nlen], addt)
	}
}

func subsets(nums []int) [][]int {
	res = [][]int{}
	dfs(nums, []int{})
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}
