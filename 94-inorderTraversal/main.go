package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func help(root *TreeNode) {
	if root.Left != nil {
		help(root.Left)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		help(root.Right)
	}
}
func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	if root == nil {
		return res
	}
	help(root)
	return res
}

func main() {
	fmt.Println(inorderTraversal(&TreeNode{0, nil, nil}))
}
