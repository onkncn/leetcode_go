package main

import "fmt"

var flag map[int]bool

func integerReplacement(n int) int {
	dp := [][]int{{n}}
	index := 0
	flag = make(map[int]bool)
	for true {
		tres := []int{}
		for _, n := range dp[index] {
			if n == 1 {
				return index
			}
			if n%2 == 0 {
				if !flag[n/2] {
					tres = append(tres, n/2)
					flag[n/2] = true
				}
			} else {
				if !flag[n+1] {
					tres = append(tres, n+1)
					flag[n+1] = true
				}
				if !flag[n-1] {
					tres = append(tres, n-1)
					flag[n-1] = true
				}
			}
		}
		index++
		fmt.Println(tres)
		dp = append(dp, tres)
	}
	return -1
}

func main() {
	fmt.Println(integerReplacement(65535))
}

/*
	0:1
	1:2
	2:3 4                     						11 100       					2
	3:5 6 8					  						101-isSymmetric 110 1000 					2.5
	4:7 9 10 12 16									111 1001 1010 1100 10000 		3
	5:11 13 14 17 18 20 24 32						100000   						3.5
	6:15 19 21 22 23 25 26 28 31 34 36 40 48 64 	1111	11111 	1000000			4
*/
