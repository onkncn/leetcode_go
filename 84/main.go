package main

import "fmt"

func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	heights = append(heights, 0)
	maxs := 0
	mymax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Println(stack)
	for i := 0; i < len(heights); i++ {
		index := len(stack) - 1
		for index > 0 && heights[stack[index]] > heights[i] {
			//fmt.Println(stack,maxs,(i-1-stack[index-1])*heights[stack[index]])
			maxs = mymax(maxs, (i-1-stack[index-1])*heights[stack[index]])
			stack = stack[:index]
			index--
		}
		stack = append(stack, i)
	}
	return maxs
}

func main() {
	largestRectangleArea([]int{3, 3, 2, 3, 3})
}
