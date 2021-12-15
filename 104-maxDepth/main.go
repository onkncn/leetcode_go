package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 1
	var dfs func(root *TreeNode, d int)
	dfs = func(root *TreeNode, d int) {
		if max < d {
			max = d
		}
		if root.Left != nil {
			dfs(root.Left, d+1)
		}
		if root.Right != nil {
			dfs(root.Right, d+1)
		}
	}
	dfs(root, 1)
	return max
}

func main() {

}
