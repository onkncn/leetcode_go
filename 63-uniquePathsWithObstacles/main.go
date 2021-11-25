package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	f[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j > 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[m-1]
}

func main() {
}
