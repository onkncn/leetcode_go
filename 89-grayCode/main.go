package main

import "fmt"

func reverse(b []int) []int {
	lenb := len(b)
	t := make([]int, lenb)
	for i := 0; i < lenb; i++ {
		t[i] = b[lenb-1-i]
	}
	return t
}

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}
	res := grayCode(n - 1)
	lr := len(res)
	res = append(res, reverse(res)...)
	for i := 0; i < lr; i++ {
		res[i] = res[i] << 1
	}
	for i := lr; i < 2*lr; i++ {
		res[i] = res[i]<<1 + 1
	}
	return res
}

func main() {
	fmt.Println(grayCode(2))
}
