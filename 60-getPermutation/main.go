package main

import "fmt"

var index int
var nums []int
var flag []bool

func getPermutation(n int, k int) string {
	ft := make([]int, n+1)
	ft[0] = 1
	for i := 1; i <= n; i++ {
		ft[i] = i * ft[i-1]
	}
	k--
	flag := make([]bool, n+1)
	s := ""
	times := n
	for times > 0 {
		t := k/ft[times-1] + 1
		//fmt.Println("t:  ",t)
		i := 1
		for ; i <= n; i++ {
			if !flag[i] {
				t--
			}
			if t == 0 {
				break
			}
		}
		flag[i] = true
		s += string(i + '0')
		k %= ft[times-1]
		times--
	}
	return s
}

func dfs(n, k, has int) bool {
	if has == n {
		fmt.Println(nums, "  ", index)
		index++
		if index == k {
			return true
		}
	} else {
		for i := 1; i <= n; i++ {
			if !flag[i] {
				flag[i] = true
				nums[has] = i
				if dfs(n, k, has+1) {
					return true
				}
				flag[i] = false
			}
		}
	}
	return false
}
func getPermutation2(n int, k int) string {
	nums = make([]int, n)
	flag = make([]bool, n+1)
	index = 0
	dfs(n, k, 0)
	s := ""
	for _, c := range nums {
		s = s + string(c+'0')
	}
	return s
}

func main() {
	fmt.Println(getPermutation2(5, 100))
	fmt.Println(getPermutation(5, 100))
}
