package main

func findRepeatedDnaSequences(s string) []string {
	l := len(s)
	if l < 11 {
		return []string{}
	}
	mp := make(map[string]int)
	res := []string{}
	for i := 10; i <= l; i++ {
		mp[s[i-10:i]]++
		if mp[s[i-10:i]] == 2 {
			res = append(res, s[i-10:i])
		}
	}
	return res
}

func main() {

}
