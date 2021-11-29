package main

import (
	"fmt"
)

func addflag(flag_all, c int) int {
	if (flag_all>>c)%2 == 0 {
		flag_all += 1 << c
	}
	return flag_all
}
func exflag(flag_all, c int) int {
	return flag_all - 1<<c
}

func main() {
	arr := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr == arr2)
	fmt.Println(arr[0:1])
}

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	//1.选基准，定义左右指针
	s, e := 0, len(arr)-1
	flag := arr[0]

	//定义i从1开始，并且小于右边
	for i := 1; i <= e; {
		if arr[i] > flag {
			//如果大于flag，则和右边交换，并且右指针减1
			arr[i], arr[e] = arr[e], arr[i]
			e--
		} else {
			//否则和左边交换
			arr[i], arr[s] = arr[s], arr[i]
			i++
			s++
		}
	}
	//递归开始到左边，左边+1到截至
	go QuickSort(arr[:s])
	QuickSort(arr[s+1:])
}
