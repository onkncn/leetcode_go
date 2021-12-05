package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mymax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func mymin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func help(root *TreeNode, l, r int) (bool, int, int) {
	if root == nil {
		return true, l, r
	}
	if root.Left != nil {
		tl := mymin(l, root.Val)
		le, ll, lr := help(root.Left, tl, r)
		if !le {
			return false, 0, 0
		}
	}
}

func isValidBST(root *TreeNode) bool {
	e, _, _ := help(root, root.Val, root.Val)
	return e
}
func main() {

}
