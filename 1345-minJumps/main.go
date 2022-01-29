package main

import "fmt"

func minJumps(arr []int) int {
	l := len(arr)
	if l < 2 {
		return 0
	}
	mp := map[int][]int{}
	viewed := make(map[int]bool)
	viewed_num := make(map[int]bool)
	for i, v := range arr {
		mp[v] = append(mp[v], i)
	}
	times := 0
	queue := []int{0}
	lq := 1
	for lq != 0 {
		for i := 0; i < lq; i++ {
			if queue[i] == l-1 {
				return times
			}
			if viewed_num[arr[queue[i]]] {
				continue
			}
			viewed[queue[i]] = true

			if queue[i]-1 > 0 && !viewed[queue[i]-1] {
				queue = append(queue, queue[i]-1)
			}
			if queue[i]+1 < l && !viewed[queue[i]+1] {
				queue = append(queue, queue[i]+1)
			}
			if !viewed_num[arr[queue[i]]] {
				for _, v := range mp[arr[queue[i]]] {
					if !viewed[v] {
						queue = append(queue, v)
					}
				}
				viewed_num[arr[queue[i]]] = true
			}
		}
		queue = queue[lq:]
		lq = len(queue)
		//fmt.Println(queue)
		times++
	}
	return -1
}

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}
