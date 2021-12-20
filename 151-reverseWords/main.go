package main

func reverseWords(s string) string {
	res := ""
	word := ""
	words := []string{}
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	res = words[len(words)-1]
	for i := len(words) - 2; i >= 0; i-- {
		res += " " + words[i]
	}
	return res
}

func main() {

}
