package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(inorder)
	if l == 0 {
		return nil
	}
	root := &TreeNode{postorder[l-1], nil, nil}
	i := 0
	for i = 0; i < l; i++ {
		if inorder[i] == postorder[l-1] {
			break
		}
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:l-1])
	return root
}

func main() {

}
