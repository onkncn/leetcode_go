package main

import "fmt"

func cal(s ,p byte)bool{
	if s==p||p=='*'||p=='?'{
		return true
	}
	return false
}
func isMatch(s string, p string) bool {
	ls := len(s)
	lp := len(p)
	if ls == 0 {
		for i := 0; i < lp; i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}
	if lp == 0 {
		return false
	}
	dp := make([][]bool, ls)

	for i := 0; i < ls; i++ {
		dp[i] = make([]bool, lp)
	}
	dp[0][0] = cal(s[0], p[0])
	if p[0]=='?'||p[0]=='*'{
		for i:=0;i<ls;i++{
			dp[i][0]=true
		}
	}
	for i:=1;i<lp;i++{
		if dp[0][i-1]&&p[i]=='*'{
			dp[0][i]=true
		}
	}
	for i := 1; i < ls; i++ {
		for j := 1; j < ls; j++ {
			if p[j-1] =='*' {
				dp[i][j]=
			}else if p[j-1] == '?'{
				dp[i][j]=
			}else{
				dp[i][j]=dp[i-1][j-1]&&cal(s[i],p[j])
			}
		}
	}
	return false
}
func main() {
	fmt.Println(isMatch())
}
