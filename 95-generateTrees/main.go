package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func add(t, node *TreeNode) *TreeNode {
	if t==nil||t.Val==0{
		return node
	}
	p := node
	tp := p
	for p != nil {
		tp = p
		if p.Val > t.Val {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if tp.Val > t.Val {
		tp.Left = t
	} else {
		tp.Right = t
	}
	return node
}

func copy(node *TreeNode) *TreeNode {
	t := TreeNode{node.Val, nil, nil}
	t.Val = node.Val
	if node.Left != nil {
		t.Left = copy(node.Left)

	}
	if node.Right != nil {
		t.Right = copy(node.Right)
	}
	return &t
}
*/

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(m, n int) []*TreeNode {
	if m > n {
		return []*TreeNode{{}}
	}
	if n == m {
		return []*TreeNode{&TreeNode{m, nil, nil}}
	}
	res := []*TreeNode{}
	for i := m; i <= n; i++ {
		left := generate(m, i-1)
		right := generate(i+1, n)
		for _, l := range left {
			for _, r := range right {
				t := &TreeNode{i, l, r}
				res = append(res, t)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(generateTrees(2))
}

/*
1
2
5
14
*/
