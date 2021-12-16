package main

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] == 'X' || board[x][y] == 'T' {
			return
		}
		board[x][y] = 'T'
		for i := 0; i < 4; i++ {
			dfs(x+direction[i][0], y+direction[i][1])
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

func main() {

}
