package main

import "fmt"

func numTrees(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	if n == 2 {
		return 2
	}
	sum := 0
	for i := 1; i <= n; i++ {
		l := numTrees(i - 1)
		r := numTrees(n - i)
		sum += l * r
	}
	return sum
}

func CTL(n int) int {
	C := 1
	for i := 1; i <= n; i++ {
		C = C * (4*i - 2) / (i + 1)
	}
	return C
}

func main() {
	for i := 1; i <= 19; i++ {
		fmt.Printf("else if n==%d{\n", i)
		fmt.Printf("return %d\n", numTrees(i))
		fmt.Print("}")
	}
}
