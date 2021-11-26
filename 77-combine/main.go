package main

import "fmt"

var res [][]int
var nn, kk int

func dfs(c int, t []int) {
	if c == kk {
		res = append(res, t)
		return
	}
	i := 1
	if len(t) != 0 {
		i = t[len(t)-1] + 1
	}
	for ; i <= nn-kk+c+1; i++ {
		adr := make([]int, len(t)+1)
		copy(adr, t)
		adr[len(t)] = i
		fmt.Println(adr)
		dfs(c+1, adr)
	}
}

func combine(n int, k int) [][]int {
	if n < k {
		return res
	}
	nn, kk = n, k
	dfs(0, []int{})
	return res
}

func main() {
	fmt.Println(combine(3, 3))
}
