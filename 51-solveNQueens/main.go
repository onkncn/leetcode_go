package main

import (
	"fmt"
)

var c, r, x, y []bool
var mapq [][]byte
var times int
var res [][]string

func mycopy(a [][]byte) []string {
	l := len(a)
	res := make([]string, l)
	for i := 0; i < l; i++ {
		res[i] = string(a[i])
	}
	return res
}
func dfs(i, j, n, k int) {
	if c[i] || r[j] || x[i+j] || y[i-j+n-1] {
		return
	}
	if k == 0 {
		mapq[i][j] = 'Q'
		res = append(res, mycopy(mapq))
		mapq[i][j] = '.'
		return
	}
	//fmt.Println(n-k, i, j)
	c[i], r[j], x[i+j], y[i-j+n-1] = true, true, true, true
	mapq[i][j] = 'Q'
	for t := i*n + j; t < n*n; t++ {
		dfs(t/n, t%n, n, k-1)
	}
	c[i], r[j], x[i+j], y[i-j+n-1] = false, false, false, false
	mapq[i][j] = '.'
}

func solveNQueens(n int) [][]string {
	res = [][]string{}
	c, r, x, y = make([]bool, n), make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	mapq = make([][]byte, n)
	for i := 0; i < n; i++ {
		mapq[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			mapq[i][j] = '.'
		}
	}
	times = 0
	for i := 0; i < n; i++ {
		dfs(0, i, n, n-1)
	}
	return res
}

func main() {
	fmt.Println(solveNQueens(4))
}
