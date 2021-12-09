package main

import "fmt"

func validTicTacToe(board []string) bool {
	nx, no := 0, 0
	c, r, x := [3]int{}, [3]int{}, [2]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				c[i] = c[i] - 1<<j
				r[j] = r[j] - 1<<i
				if i == j {
					x[0] = x[0] - 1<<j
				}
				if i+j == 2 {
					x[1] = x[1] - 1<<j
				}
				nx++
			} else if board[i][j] == 'O' {
				c[i] = c[i] + 1<<j
				r[j] = r[j] + 1<<i
				if i == j {
					x[0] = x[0] + 1<<j
				}
				if i+j == 2 {
					x[1] = x[1] + 1<<j
				}
				no++
			}
		}
	}
	if no > nx || nx > no+1 {
		return false
	}
	flagcx, flagco := false, false
	for i := 0; i < 3; i++ {
		if c[i] == 7 {
			if !flagco {
				flagco = true
			} else {
				return false
			}
		}
		if c[i] == -7 {
			if !flagcx {
				flagcx = true
			} else {
				return false
			}
		}
	}
	flagrx, flagro := false, false
	for i := 0; i < 3; i++ {
		if r[i] == 7 {
			if !flagro {
				flagro = true
			} else {
				return false
			}
		}
		if r[i] == -7 {
			if !flagrx {
				flagrx = true
			} else {
				return false
			}
		}
	}
	if (flagrx || flagcx) && (flagro || flagco) {
		return false
	}
	if nx == no && (flagrx || flagcx || x[0] == -7 || x[1] == -7) {
		return false
	}
	if nx == no+1 && (flagro || flagco || x[0] == 7 || x[1] == 7) {
		return false
	}
	return true
}

func main() {
	fmt.Println(1 << 2)
	fmt.Println(validTicTacToe([]string{"XXO", "XOX", "OXO"}))
}
