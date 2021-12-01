package main

import (
	"fmt"
	"math/rand"
)

func qsort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	r := rand.Intn(l)
	m := nums[r]
	nums[r], nums[0] = nums[0], nums[r]
	h, t := 1, l-1
	for h <= t {
		if nums[h] < m {
			nums[h], nums[h-1] = nums[h-1], nums[h]
			h++
		} else {
			nums[t], nums[h] = nums[h], nums[t]
			t--
		}
	}
	qsort(nums[:t])
	qsort(nums[h:])
	return nums
}

func mergesort(a []int, l, r int) []int {
	if l+1 == r {
		return a
	}
	m := (l + r) / 2
	mergesort(a, l, m)
	mergesort(a, m, r)
	return merge(a, l, m, r)
}
func merge(a []int, l, m, r int) []int {
	t := make([]int, r-l)
	p1, p2, i := l, m, 0
	for p1 < m && p2 < r {
		if a[p1] < a[p2] {
			t[i] = a[p1]
			p1++
		} else {
			t[i] = a[p2]
			p2++
		}
		i++
	}
	for p1 < m {
		t[i] = a[p1]
		p1++
		i++
	}
	for p2 < r {
		t[i] = a[p2]
		p2++
		i++
	}
	for i = l; i < r; i++ {
		a[i] = t[i-l]
	}
	return a
}

func main() {
	numbers := []int{5, 2, 3, 1, 4}
	mergesort(numbers, 0, len(numbers))
	fmt.Println(numbers)
}
