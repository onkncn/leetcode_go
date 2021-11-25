package main

import "fmt"

func generateMatrix(n int) [][]int {
	var res [][]int
	for i := 0; i < n; i++ {
		var l []int
		for j := 0; j < n; j++ {
			l = append(l, 0)
		}
		res = append(res, l)
	}
	index := 1
	for i := 0; i < (n+1)/2; i++ {
		//right
		for j := i; j < n-i-1; j++ {
			res[i][j] = index
			index++
		}
		//down
		for j := i; j < n-i-1; j++ {
			res[j][n-1-i] = index
			index++
		}
		//left
		for j := i; j < n-i-1; j++ {
			res[n-1-i][n-1-j] = index
			index++
		}
		//up
		for j := i; j < n-i-1; j++ {
			res[n-1-j][i] = index
			index++
		}
	}
	if n%2 == 1 {
		res[n/2][n/2] = n * n
	}
	return res
}
func main() {
	fmt.Println(generateMatrix(5))
}

/*
1

1 2
4 3

1 2 3
8 9 4
7 6 5

1 1 1 1
1 1 1 1
1 1 1 1
1 1 1 1


1  2  3  4  5
16 17 18 19 6
15 24 25 20 7
14 23 22 21 8
13 12 11 10 9


*/
