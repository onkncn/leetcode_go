package main

func countBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'X' {
			return
		}
		board[i][j] = '.'
		for t := 0; t < 4; t++ {
			dfs(i+direction[t][0], j+direction[t][1])
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				res++
				//fmt.Println(i,j)
				dfs(i, j)
			}
		}
	}
	return res
}

func main() {

}
