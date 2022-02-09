/**
 * @Author Oliver
 * @Date 2/8/22
 **/

package main

import "fmt"

var direction = [9][2]int{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, 0}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}

func gridIllumination(n int, lamps [][]int, queries [][]int) (res []int) {
	lampsMap := map[[2]int]bool{}
	c := map[int]int{}
	r := map[int]int{}
	x := map[int]int{}
	y := map[int]int{}
	var operate = func(i, j, cal int) {
		if cal == 1 {
			if lampsMap[[2]int{i, j}] {
				return
			}
			lampsMap[[2]int{i, j}] = true
		} else {
			lampsMap[[2]int{i, j}] = false
		}
		c[i] += cal
		r[j] += cal
		x[i+j] += cal
		y[i-j] += cal
	}
	for i := 0; i < len(lamps); i++ {
		operate(lamps[i][0], lamps[i][1], 1)
	}
	for i := 0; i < len(queries); i++ {
		//fmt.Println(c[queries[i][0]], r[queries[i][1]], x[queries[i][0]+queries[i][1]], y[queries[i][0]-queries[i][1]], lampsMap)
		if c[queries[i][0]] > 0 || r[queries[i][1]] > 0 || x[queries[i][0]+queries[i][1]] > 0 || y[queries[i][0]-queries[i][1]] > 0 {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
		for j := 0; j < 9; j++ {
			opx, opy := queries[i][0]+direction[j][0], queries[i][1]+direction[j][1]
			if lampsMap[[2]int{opx, opy}] {
				operate(opx, opy, -1)
			}
		}
	}
	return res
}

func main() {
	var (
		n = 6
		l = [][]int{{2, 5}, {4, 2}, {0, 3}, {0, 5}, {1, 4}, {4, 2}, {3, 3}, {1, 0}}
		q = [][]int{{4, 3}, {3, 1}, {5, 3}, {0, 5}, {4, 4}, {3, 3}}
	)
	fmt.Println(gridIllumination(n, l, q))
}
