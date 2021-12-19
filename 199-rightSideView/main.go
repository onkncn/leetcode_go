package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	right := []int{}
	var dfs func(root *TreeNode, deep int)
	dfs = func(root *TreeNode, deep int) {
		if root != nil {
			if deep == len(right) {
				right = append(right, root.Val)
			}
			dfs(root.Right, deep+1)
			dfs(root.Left, deep+1)
		}
	}
	dfs(root, 0)
	return right
}

func main() {

}
