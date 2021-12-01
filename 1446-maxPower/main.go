package main

func maxPower(s string) int {
	max := 1
	alen := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			alen++
		} else {
			if alen > max {
				max = alen
			}
			alen = 1
		}
	}
	if alen > max {
		max = alen
	}
	return max
}
func main() {

}
