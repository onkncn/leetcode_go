package main

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	for i := 1; i < l; i++ {
		line := len(triangle[i])
		triangle[i][0] += triangle[i-1][0]
		triangle[i][line-1] += triangle[i-1][line-2]
		for j := 1; j < line-1; j++ {
			triangle[i][j] += mymin(triangle[i-1][j], triangle[i-1][j-1])
		}
	}
	min := triangle[l-1][0]
	for i := 1; i < len(triangle[l-1]); i++ {
		if triangle[l-1][i] < min {
			min = triangle[l-1][i]
		}
	}
	return min
}
func mymin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
