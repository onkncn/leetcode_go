package main

import (
	"container/heap"
)

type hp struct {
	nums         [][2]int
	nums1, nums2 []int
}

func (h hp) Len() int      { return len(h.nums) }
func (h hp) Swap(i, j int) { h.nums[i], h.nums[j] = h.nums[j], h.nums[i] }
func (h hp) Less(i, j int) bool {
	return h.nums1[h.nums[i][0]]+h.nums2[h.nums[i][1]] < h.nums1[h.nums[j][0]]+h.nums2[h.nums[j][1]]
}
func (h *hp) Push(i interface{}) { h.nums = append(h.nums, i.([2]int)) }
func (h *hp) Pop() interface{} {
	old := h.nums
	l := len(h.nums)
	h.nums = h.nums[0 : l-1]
	return old[l-1]
}
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := &hp{[][2]int{}, nums1, nums2}
	for i := 0; i < n; i++ {
		heap.Push(h, [2]int{0, i})
	}
	//fmt.Println(h)
	ans = [][]int{}
	for i := 0; i < k; i++ {
		if len(h.nums) == 0 {
			return ans
		}
		t := []int{h.nums[0][0], h.nums[0][1]}
		heap.Pop(h)
		if t[0]+1 < m {
			heap.Push(h, [2]int{t[0] + 1, t[1]})
		}
		ans = append(ans, []int{nums1[t[0]], nums2[t[1]]})
		//fmt.Println(h)
	}
	return ans
}

func main() {

}
