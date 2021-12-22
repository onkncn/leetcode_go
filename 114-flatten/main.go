package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var help func(root *TreeNode) *TreeNode
	help = func(root *TreeNode) *TreeNode {
		if root.Left == nil && root.Right == nil {
			return root
		}
		l, r, last := root.Left, root.Right, root
		if root.Left != nil {
			last = help(root.Left)
		}
		root.Left = nil
		root.Right = l
		last.Right = r
		if root.Right != nil {
			last = help(root.Right)
		}
		return last
	}
	help(root)
}

func main() {

}
