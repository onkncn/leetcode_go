package main

import "fmt"

var flag [12][12]bool
var res []int
var mix [][]int

func next(i, j, d int) []int {
	//fmt.Printf("%d  %d\n",i,j)
	if flag[i+1][j+1] {
		return res
	} else {
		res = append(res, mix[i][j])
		flag[i+1][j+1] = true
	}
	switch {
	case d == 0:
		{
			if !flag[i+1][j+2] {
				return next(i, j+1, 0)
			} else {
				return next(i+1, j, 1)
			}
		}
	case d == 1:
		{
			if !flag[i+2][j+1] {
				return next(i+1, j, 1)
			} else {
				return next(i, j-1, 2)
			}
		}
	case d == 2:
		{
			if !flag[i+1][j] {
				return next(i, j-1, 2)
			} else {
				return next(i-1, j, 3)
			}
		}
	case d == 3:
		{
			if !flag[i][j+1] {
				return next(i-1, j, 3)
			} else {
				return next(i, j+1, 0)
			}
		}
	default:
		{
			if !flag[i+1][j+2] {
				return next(i, j+1, 0)
			} else {
				return next(i+1, j, 1)
			}
		}

	}
	return res
}

func spiralOrder(matrix [][]int) []int {
	res = []int{}
	flag = [12][12]bool{}
	mlen := len(matrix)
	nlen := len(matrix[0])
	mix = matrix
	for i := 0; i <= mlen; i++ {
		flag[i][0] = true
		flag[i][nlen+1] = true
	}
	for i := 0; i <= nlen; i++ {
		flag[0][i] = true
		flag[mlen+1][i] = true
	}
	next(0, 0, 0)
	return res
}

func main() {
	mix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(mix))
}

/*
t t t t  t
t 4 8 12 t
t 3 7 11 t
t 2 6 10 t
t 1 5 9  t
t t t t  t
*/
