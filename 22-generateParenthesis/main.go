package main

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, q := range dp[j] {
				for _, p := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+p+")"+q)
				}
			}
		}
	}
	return dp[n]
}

func main() {

}
