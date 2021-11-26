package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum = 0

func abse(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func findTiltbfs(root *TreeNode) int {
	r := 0
	l := 0
	if root.Right != nil {
		r = root.Right.Val
		r += findTiltbfs(root.Right)
	}
	if root.Left != nil {
		l = root.Left.Val
		l += findTiltbfs(root.Left)
	}
	sum += abse(r, l)
	return r + l
}
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum = 0
	findTiltbfs(root)
	return sum
}
func main() {
	fmt.Println("")
}
