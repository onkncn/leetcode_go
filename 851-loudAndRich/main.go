package main

func loudAndRich(richer [][]int, quiet []int) []int {
	lq := len(quiet)
	if lq < 1 {
		return []int{0}
	}
	res := make([]int, lq)
	flag := make(map[int]bool)
	var dfs func(r int)
	for j := 0; j < lq; j++ {
		flag = make(map[int]bool)
		res[j] = j
		dfs = func(r int) {
			if flag[r] {
				return
			}
			flag[r] = true
			for i := 0; i < len(richer); i++ {
				if richer[i][1] == r {
					if quiet[richer[i][0]] < quiet[res[j]] {
						res[j] = richer[i][0]
					}
					if richer[i][0] < j {
						if quiet[res[j]] > quiet[res[richer[i][0]]] {
							res[j] = res[richer[i][0]]
						}
					} else {
						dfs(richer[i][0])
					}
				}
			}
		}
		dfs(j)
	}
	return res
}

func main() {
}
