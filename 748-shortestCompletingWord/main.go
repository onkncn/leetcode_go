package main

func shortestCompletingWord(licensePlate string, words []string) string {
	var flag = [26]int{}
	ls := ""
	for _, c := range licensePlate {
		if c >= 'a' && c <= 'z' {
			flag[c-'a']++
		}
		if c >= 'A' && c <= 'Z' {
			flag[c-'A']++
		}
	}
	for _, s := range words {
		var t = [26]int{}
		for _, c := range s {
			if len(s) >= len(ls) && ls != "" {
				continue
			}
			if c >= 'a' && c <= 'z' {
				if t[c-'a'] < flag[c-'a'] {
					t[c-'a']++
				}
			}
			if c >= 'A' && c <= 'Z' {
				if t[c-'A'] < flag[c-'A'] {
					t[c-'A']++
				}
			}
			if t == flag {
				ls = s
			}
		}
	}
	return ls
}
func main() {

}
