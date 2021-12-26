package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack_nums := []int{}
	l := 0
	for i := 0; i < len(tokens); i++ {
		if tokens[i][0] >= '0' && tokens[i][0] <= '9' || tokens[i][0] == '-' && len(tokens[i]) != 1 {
			s, _ := strconv.Atoi(tokens[i])
			stack_nums = append(stack_nums, s)
			l++
		} else {
			if tokens[i][0] == '+' {
				stack_nums[l-2] = stack_nums[l-2] + stack_nums[l-1]
			} else if tokens[i][0] == '-' {
				stack_nums[l-2] = stack_nums[l-2] - stack_nums[l-1]
			} else if tokens[i][0] == '*' {
				stack_nums[l-2] = stack_nums[l-2] * stack_nums[l-1]
			} else if tokens[i][0] == '/' {
				stack_nums[l-2] = stack_nums[l-2] / stack_nums[l-1]
			}
			stack_nums = stack_nums[:l-1]
			l--
		}
	}
	return stack_nums[0]
}

func main() {

}
