package main

func lengthOfLastWord(s string) int {
	llen := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if llen != 0 {
				res = llen
			}
			llen = 0
		} else {
			llen++
		}
	}
	if llen == 0 {
		return res
	}
	return llen
}

func main() {

}
