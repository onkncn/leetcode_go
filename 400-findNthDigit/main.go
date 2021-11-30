package main

import "fmt"

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	sub_d := 9
	sub_n := 1
	for n-sub_d*sub_n > 0 {
		n -= sub_d * sub_n
		sub_d *= 10
		sub_n++
	}
	c := (n - 1) % sub_n       //第几位.   1
	p := (n-1)/sub_n + sub_d/9 //第多少个数 1
	//fmt.Printf("P:%d  C:%d   ",p,c)
	for i := 1; i < sub_n-c; i++ {
		p /= 10
	}
	return p % 10
}

func main() {
	for i := 1; i <= 200; i++ {
		fmt.Println(findNthDigit(i))
	}
}

// 1 2 3 4 5 6 7 8 9 10 11 12
/*
   9.    *1.
   90.   *2
   900.  *3
   9000. *4
   90000.*5
*/
