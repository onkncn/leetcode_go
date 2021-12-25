package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	ss := strings.Split(text, " ")
	res := []string{}
	for i := 1; i < len(ss)-1; i++ {
		if ss[i-1] == first && ss[i] == second {
			res = append(res, ss[i+1])
		}
	}
	return res
}

func main() {

}
