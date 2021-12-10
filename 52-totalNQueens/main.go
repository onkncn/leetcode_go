package main

var c, r, x, y []bool
var res [][]string
var times int

func dfs(i, j, n, k int) {
	if c[i] || r[j] || x[i+j] || y[i-j+n-1] {
		return
	}
	if k == 0 {
		times++
		return
	}
	//fmt.Println(n-k,i,j)
	c[i], r[j], x[i+j], y[i-j+n-1] = true, true, true, true
	for t := i*n + j; t < n*n; t++ {
		dfs(t/n, t%n, n, k-1)
	}
	c[i], r[j], x[i+j], y[i-j+n-1] = false, false, false, false
}

func totalNQueens(n int) int {
	c, r, x, y = make([]bool, n), make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	times = 0
	for i := 0; i < n; i++ {
		dfs(0, i, n, n-1)
	}
	return times
}

func main() {

}
