package main

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
func maxPathSum(root *TreeNode) int {
	var most func(root *TreeNode) int
	max := -1111
	most = func(root *TreeNode) int {
		if root == nil {
			return -1111
		}
		l, r, m := most(root.Left), most(root.Right), root.Val
		max = mymax(mymax(mymax(mymax(mymax(mymax(l, r), m), l+m), r+m), l+r+m), max)
		return mymax(mymax(l, r), 0) + m
	}
	most(root)
	return max
}
func main() {

}
