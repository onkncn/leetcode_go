package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	maxi, maxj := make([]int, n), make([]int, n)
	mymax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mymin := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maxi[j] = mymax(grid[i][j], maxi[j])
			maxj[i] = mymax(grid[i][j], maxj[i])
		}
	}
	//fmt.Println(maxi,maxj)
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += mymin(maxj[i], maxi[j]) - grid[i][j]
		}
	}
	return sum
}

func main() {

}
