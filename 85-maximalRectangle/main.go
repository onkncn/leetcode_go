package main

import "fmt"

func mymax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	heights = append(heights, 0)
	maxs := 0
	for i := 0; i < len(heights); i++ {
		index := len(stack) - 1
		for index > 0 && heights[stack[index]] > heights[i] {
			maxs = mymax(maxs, (i-1-stack[index-1])*heights[stack[index]])
			stack = stack[:index]
			index--
		}
		stack = append(stack, i)
	}
	return maxs
}
func maximalRectangle(matrix [][]byte) int {
	m, n, maxs := len(matrix), len(matrix[0]), 0
	tmatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		tmatrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			tmatrix[0][i] = 1
		}
	}
	maxs = mymax(maxs, largestRectangleArea(tmatrix[0]))
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				tmatrix[i][j] = tmatrix[i-1][j] + 1
			}
		}
		maxs = mymax(maxs, largestRectangleArea(tmatrix[i]))
	}
	return maxs
}

func main() {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}
