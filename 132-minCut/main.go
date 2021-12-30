package main

func minCut(s string) int {
	l := len(s)
	isR := make([][]bool, l)
	for i := range isR {
		isR[i] = make([]bool, l)
		for j := range isR[i] {
			isR[i][j] = true
		}
	}
	for i := l - 1; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			isR[i][j] = s[i] == s[j] && isR[i+1][j-1]
		}
	}
	// for t:=range isR{
	//     fmt.Println(isR[t])
	// }
	dp := make([]int, l)
	for i := 1; i < l; i++ {
		if isR[0][i] {
			dp[i] = 0
			continue
		}
		dp[i] = dp[i-1] + 1
		for j := 1; j < i; j++ {
			//fmt.Println(j,i,isR[j][i])
			if isR[j][i] && dp[j-1]+1 < dp[i] {
				dp[i] = dp[j-1] + 1
			}
		}
	}
	//fmt.Println(dp)
	return dp[l-1]
}

func main() {
	minCut("123321")
}
