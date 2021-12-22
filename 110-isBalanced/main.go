package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl := height(root.Left)
	if hl == -1 {
		return -1
	}
	hr := height(root.Right)
	if hl-hr > 1 || hr-hl > 1 || hr == -1 {
		return -1
	}
	return 1 + mymax(hl, hr)
}
func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func main() {

}
