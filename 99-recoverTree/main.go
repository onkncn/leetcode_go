package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pre int
var flag bool
var error [2]*TreeNode
var index_err int

func help(root *TreeNode) {
	if root.Left != nil {
		help(root.Left)
	}
	if !flag {
		flag = true
	} else {
		if root.Val <= pre {
			if index_err == 0 {
				index_err++
				error[index_err] = root
			} else {
				error[1] = root
				return
			}
		}
	}
	pre = root.Val
	if index_err == 0 {
		error[index_err] = root
	}
	if root.Right != nil {
		help(root.Right)
	}
}
func recoverTree(root *TreeNode) {
	flag = false
	index_err = 0
	help(root)
	//fmt.Println(error[0].Val,error[1].Val)
	error[0].Val, error[1].Val = error[1].Val, error[0].Val
}

func main() {

}
