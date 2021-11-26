package main

import "fmt"

var flag [][]bool
var ans string
var n, m int
var wmap [][]byte
var dir = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func dfs(x, y int, word string) bool {
	if word != ans[0:len(word)] {
		return false
	}
	if word == ans {
		return true
	}
	if x >= n || x < 0 || y >= m || y < 0 || flag[x][y] {
		return false
	}
	flag[x][y] = true
	for i := 0; i < 4; i++ {
		if dfs(x+dir[i][0], y+dir[i][1], word+string(wmap[x][y])) {
			return true
		}
	}
	flag[x][y] = false
	return false
}
func exist(board [][]byte, word string) bool {
	ans = word
	wmap = board
	n, m = len(board), len(board[0])
	flag = make([][]bool, n)
	for i, _ := range flag {
		flag[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, "") {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}
	word := "abcdeaafg"
	fmt.Println(exist(board, word))

	fmt.Println(exist([][]byte{{'a'}}, ""))
}
