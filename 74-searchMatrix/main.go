package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	c := sort.Search(n*m, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})
	return c < n*m && matrix[c/n][c%n] == target
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i <= 10; i++ {
		fmt.Println(searchMatrix(matrix, i))
	}
}
