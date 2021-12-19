package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	var sort func(a []int, s int) []int
	sort = func(a []int, s int) []int {
		l := len(a)
		if l < 2 {
			return a
		}
		r := rand.Intn(l)
		a[r], a[0] = a[0], a[r]
		h, t, m := 1, l-1, a[0]
		for h <= t {
			if a[h] < m {
				a[h], a[h-1] = a[h-1], a[h]
				h++
			} else {
				a[h], a[t] = a[t], a[h]
				t--
			}
		}
		if t+s > len(nums)-k {
			sort(a[:t], s)
		} else {
			sort(a[h:], s+h)
		}
		return a
	}
	sort(nums, 0)
	fmt.Println(nums)
	return nums[len(nums)-k]
}
func main() {

}
