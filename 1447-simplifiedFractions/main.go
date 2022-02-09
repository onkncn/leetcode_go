package main

import "strconv"

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
func simplifiedFractions(n int) (res []string) {
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				res = append(res, strconv.Itoa(i)+"/"+strconv.Itoa(j))
			}
		}
	}
	return res
}
