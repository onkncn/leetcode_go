package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int)
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != '1' {
			return
		}
		grid[x][y] = '0'
		for i := 0; i < 4; i++ {
			dfs(x+direction[i][0], y+direction[i][1])
		}
	}
	nums := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				nums++
			}
		}
	}
	return nums
}

func main() {

}
