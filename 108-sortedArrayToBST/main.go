package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	m := l / 2
	return &TreeNode{nums[m], sortedArrayToBST(nums[:m]), sortedArrayToBST(nums[m+1:])}
}

func main() {

}
