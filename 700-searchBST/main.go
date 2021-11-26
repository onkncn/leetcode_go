package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	for root.Val != val {
		if root.Val < val {
			if root.Right == nil {
				return nil
			}
			root = root.Right
		} else {
			if root.Left == nil {
				return nil
			}
			root = root.Left
		}
	}
	return root
}

func main() {

}
