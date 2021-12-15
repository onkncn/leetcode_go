package main

import (
	"fmt"
	"math"
)

func binarySearch(c int) int {
	a := func(x int) int {
		fmt.Printf("func 3, x is %d\n", x)
		return 5
	}
	return a(3)
}
func main() {
	fmt.Println(math.Atan2(-0.01, -1) / math.Pi * 180)
}
