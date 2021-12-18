package main

func findJudge(n int, trust [][]int) int {
	l := len(trust)
	flag := make([]int, n+1)
	isTrust := make([]bool, n+1)
	for i := 0; i < l; i++ {
		flag[trust[i][1]]++
		isTrust[trust[i][0]] = true
	}
	for i := 1; i <= n; i++ {
		if flag[i] == n-1 && !isTrust[i] {
			return i
		}
	}
	return -1
}

func main() {

}
