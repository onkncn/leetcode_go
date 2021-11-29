package main

import (
	"math"
	"sort"
)

func kthSmallestPrimeFraction(arr []int, k int) []int {
	la := len(arr)
	if k == 1 {
		return []int{arr[0], arr[la-1]}
	}
	farr := make([]float64, la)
	//mk := la*(la-1)/2
	mp := map[float64][]int{}
	for i := 0; i < la; i++ {
		farr[i] = math.Log(float64(arr[i]))
	}
	ans := []float64{}
	for i := 0; i < la-1; i++ {
		for j := i + 1; j < la; j++ {
			mp[farr[i]-farr[j]] = []int{i, j}
			ans = append(ans, farr[i]-farr[j])
		}
	}
	sort.Float64s(ans)
	return []int{arr[mp[ans[k-1]][0]], arr[mp[ans[k-1]][1]]}
}

func main() {

}
