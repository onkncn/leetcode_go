package main

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
	if root.Right != nil {
		help(root.Right)
	}
	res = append(res, root.Val)
}
func postorderTraversal(root *TreeNode) []int {
	res = []int{}
	if root == nil {
		return res
	}
	help(root)
	return res
}

func main() {

}
