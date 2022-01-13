package main

import "fmt"

func catMouseGame(graph [][]int) int {
	n := len(graph)
	viewed := make(map[int]bool, n)
	qm, qc := []int{1}, []int{2}
	lm, lc := 1, 1
	step := make([]int, n)
	circle := make([]int, n)
	step[0] = 0
	var dfs func(t int) (c int)
	dfs = func(t int) (c int) {
		viewed[t] = true
		c = n + 1
		for _, i := range graph[t] {
			if i == 0 {
				continue
			}
			if viewed[i] {
				if step[t]-step[i] > 1 && step[t]-step[i] < c {
					c = step[t] - step[i] + 1
				}
			} else {
				fmt.Println(i)
				step[i] = step[t] + 1
				t := dfs(i)
				if t < c {
					c = t
				}
			}
		}
		//viewed[t]=false
		circle[t] = c
		return c
	}
	dfs(0)
	fmt.Println(circle)
	viewed[2] = true
	peace := false
	for lm != 0 && lc != 0 {
		for i := 0; i < lm; i++ {
			t := qm[i]
			if viewed[t] {
				continue
			}
			if circle[t] > 3 && circle[t] < n+1 {
				peace = true
			}
			for j := 0; j < len(graph[t]); j++ {
				if graph[t][j] == 0 {
					return 1
				}
				if !viewed[graph[t][j]] {
					qm = append(qm, graph[t][j])
				}
			}
		}
		qm = qm[lm:]
		lm = len(qm)
		if lm == 0 {
			if peace {
				return 0
			}
			return 2
		}
		for i := 0; i < lc; i++ {
			t := qc[i]
			for j := 0; j < len(graph[t]); j++ {
				if viewed[graph[t][j]] != true && graph[t][j] != 0 {
					qc = append(qc, graph[t][j])
					viewed[graph[t][j]] = true
				}
			}
		}
		qc = qc[lc:]
		lc = len(qc)
	}
	return 0
}

func main() {
	//graph:=[][]int{{3,4,6,7,9,15,16,18},{4,5,8,19},{3,4,6,9,17,18},{0,2,11,15},{0,1,10,6,2,12,14,16},{1,10,7,9,15,17,18},{0,10,4,7,9,2,11,12,13,14,15,17,19},{0,10,5,6,9,16,17},{1,9,14,15,16,19},{0,10,5,6,7,8,2,11,13,15,16,17,18},{4,5,6,7,9,18},{3,6,9,12,19},{4,6,11,15,17,19},{6,9,15,17,18,19},{4,6,8,15,19},{0,3,5,6,8,9,12,13,14,16,19},{0,4,7,8,9,15,17,18,19},{5,6,7,9,2,12,13,16},{0,10,5,9,2,13,16},{1,6,8,11,12,13,14,15,16}}
	graph := [][]int{{1, 2, 3}, {0, 2, 3}, {0, 1, 3}, {0, 2, 3}}
	catMouseGame(graph)
}
