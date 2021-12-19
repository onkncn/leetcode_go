package main

import "fmt"

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
func maxPoints(points [][]int) int {
	n := len(points)
	if n < 2 {
		return n
	}
	maxpoint := 1
	for i := 0; i < n-1; i++ {
		k := make(map[[2]int]int)
		for j := i + 1; j < n; j++ {
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				t := gcd(x, y)
				x, y = x/t, y/t
			}
			k[[2]int{x, y}]++
			if k[[2]int{x, y}] > maxpoint {
				maxpoint = k[[2]int{x, y}]
			}
		}
	}
	return maxpoint + 1
}
func main() {
	fmt.Println(maxPoints([][]int{{2, 3}, {3, 3}, {-5, 3}}))
}
