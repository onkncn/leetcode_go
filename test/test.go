package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 9, 10, 3, 15, 8, 212}
	QuickSort(arr)
	fmt.Println(arr)
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
