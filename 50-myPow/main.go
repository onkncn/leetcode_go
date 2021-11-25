package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	index := 0
	px := []float64{x}
	tn := n
	sum := float64(1)
	if n < 0 {
		tn = -n
	}
	for tn/2 != 0 {
		if tn%2 == 1 {
			sum *= px[index]
		}
		px = append(px, px[index]*px[index])
		index++
		tn /= 2
	}
	if n < 0 {
		return 1 / sum / px[index]
	}
	return sum * px[index]
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(myPow(2, -i))
	}
}
