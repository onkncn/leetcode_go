package main

import (
	"fmt"
	"strings"
)

func binarySearch(c int) int {
	a := func(x int) int {
		fmt.Printf("func 3, x is %d\n", x)
		return 5
	}
	return a(3)
}
func main() {
	a := "11111111112222222"
	fmt.Println(strings.Replace(a, "1", "3", 4))
}
