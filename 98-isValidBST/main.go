package main

import (
	"fmt"
)

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
func help(root *TreeNode) (bool, int, int) {
	if root == nil {
		return true, 0, 0
	}
	var le, re bool
	var lmax, lmin, rmax, rmin int
	vmax, vmin := root.Val, root.Val
	if root.Left != nil {
		le, lmax, lmin = help(root.Left)
		if le == false || lmax >= root.Val {
			return false, 0, 0
		}
		vmax = mymax(vmax, lmax)
		vmin = mymin(vmin, lmin)
	}
	if root.Right != nil {
		re, rmax, rmin = help(root.Right)
		if re == false || rmin <= root.Val {
			return false, 0, 0
		}
		vmax = mymax(vmax, rmax)
		vmin = mymin(vmin, rmin)
	}
	return true, vmax, vmin
}
func isValidBST_slow(root *TreeNode) bool {
	e, _, _ := help(root)
	return e
}

var pre int
var flag bool

func help2(root *TreeNode) bool {
	if root.Left != nil {
		if !help2(root.Left) {
			return false
		}
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	if root.Right != nil {
		if !help2(root.Right) {
			return false
		}
	}
	return true
}
func isValidBST(root *TreeNode) bool {
	flag = false
	pre = -1 << (32 - 1)
	return help2(root)
}
func main() {
	root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(isValidBST(root))
}
