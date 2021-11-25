package main

import (
	"math/rand"
)

type Solution struct {
	n []int
}

func Constructor(nums []int) Solution {
	return Solution{
		n: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.n
}

func (this *Solution) Shuffle() []int {
	ln := len(this.n)
	nums := make([]int, ln)
	copy(nums, this.n)
	for i := 0; i < ln; i++ {
		t := rand.Intn(i + 1)
		nums[i], nums[t] = nums[t], nums[i]
	}
	return nums
}

func main() {

}
