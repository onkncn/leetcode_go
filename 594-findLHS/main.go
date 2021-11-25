package main

import "fmt"

var nb map[int]int

func findLHS(nums []int) int {
	nb = make(map[int]int)
	sn := []int{}
	for _, n := range nums {
		if nb[n] == 0 {
			sn = append(sn, n)
		}
		nb[n]++
	}
	max := 0
	for k, v := range nb {
		if nb[k+1] != 0 && nb[k+1]+v > max {
			max = v + nb[k+1]
		}
	}
	return max
}
func main() {
	nums := []int{1, 1, 1, 1, 3, 3, 3, 3, 3, 3, 2, 2}
	fmt.Println(findLHS(nums))
}
