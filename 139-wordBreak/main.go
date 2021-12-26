package main

func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]int)
	var help func(s string, wordDict []string) bool
	help = func(s string, wordDict []string) (res bool) {
		if mp[s] == 1 {
			return true
		}
		if mp[s] == -1 {
			return false
		}
		defer func() {
			if res {
				mp[s] = 1
			} else {
				mp[s] = -1
			}
		}()
		for _, ss := range wordDict {
			if ss == s {
				return true
			}
		}
		for i := 1; i < len(s); i++ {
			if help(s[:i], wordDict) && help(s[i:], wordDict) {
				return true
			}
		}
		return false
	}
	return help(s, wordDict)
}
func main() {

}
