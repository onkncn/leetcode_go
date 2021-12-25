package main

func getc(c byte) byte {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return c
	} else if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	} else {
		return ' '
	}
}
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		c1 := getc(s[l])
		if c1 == ' ' {
			l++
			continue
		}
		c2 := getc(s[r])
		if c2 == ' ' {
			r--
			continue
		}
		if c1 != c2 {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {

}
