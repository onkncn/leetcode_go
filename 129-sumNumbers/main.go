package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	var help func(root *TreeNode, num int) bool
	help = func(root *TreeNode, num int) bool {
		if root == nil {
			return true
		}
		n := num*10 + root.Val
		flag1 := help(root.Left, n)
		flag2 := help(root.Right, n)
		if flag1 && flag2 {
			sum += n
		}
		return false
	}
	help(root, 0)
	return sum
}

func main() {

}
