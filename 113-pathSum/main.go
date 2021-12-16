package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createrTree(a []int) *TreeNode {
	l := len(a)
	if l == 0 {
		return nil
	}
	return help(a, 0)
}
func help(a []int, index int) *TreeNode {
	if a[index] == -1 {
		return nil
	}
	t := &TreeNode{a[index], nil, nil}
	if index*2+1 < len(a) {
		t.Left = help(a, index*2+1)
	}
	if index*2+2 < len(a) {
		t.Right = help(a, index*2+2)
	}
	return t
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, targetSum int, way []int) bool
	dfs = func(root *TreeNode, targetSum int, way []int) bool {
		if root == nil {
			return true
		}
		way = append(way, root.Val)
		//fmt.Println(way)
		flag1 := dfs(root.Left, targetSum-root.Val, way)
		flag2 := dfs(root.Right, targetSum-root.Val, way)
		if flag1 && flag2 {
			if targetSum-root.Val == 0 {
				tway := make([]int, len(way))
				copy(tway, way)
				res = append(res, tway)
			}
		}
		way = way[:len(way)-1]
		return false
	}
	dfs(root, targetSum, []int{})
	return res
}

func main() {
	t := createrTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, 1, 5, 1})
	fmt.Println(t)
	fmt.Println(pathSum(t, 22))
}
