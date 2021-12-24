package main

import "container/heap"

type appleDay struct {
	nums, day int
}
type heapApple []appleDay

func (h heapApple) Len() int           { return len(h) }
func (h heapApple) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h heapApple) Less(i, j int) bool { return h[i].day < h[j].day }
func (h *heapApple) Push(x interface{}) {
	ad, ok := x.(*appleDay)
	if ok {
		*h = append(*h, *ad)
	}
}
func (h *heapApple) Pop() interface{} {
	old := *h
	l := len(old)
	*h = old[:l-1]
	return old[l-1]
}

func eatenApples(apples []int, days []int) int {
	h := &heapApple{}
	heap.Init(h)
	nums := 0
	l := len(apples)
	for i := 0; i < l; i++ {
		if apples[i] != 0 {
			heap.Push(h, &appleDay{apples[i], i + days[i]})
		}
		for len(*h) > 0 && (*h)[0].day <= i {
			heap.Pop(h)
		}
		if len(*h) > 0 {
			nums++
			(*h)[0].nums--
			if (*h)[0].nums == 0 {
				heap.Pop(h)
			}
		}
	}
	now_day := l
	for len(*h) != 0 {
		if (*h)[0].day > now_day {
			if (*h)[0].nums >= (*h)[0].day-now_day {
				nums += (*h)[0].day - now_day
				now_day += (*h)[0].day - now_day
			} else {
				nums += (*h)[0].nums
				now_day += (*h)[0].nums
			}
		}
		heap.Pop(h)
	}
	return nums
}
