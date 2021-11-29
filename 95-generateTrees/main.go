package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func add(t, node TreeNode) *TreeNode {
	p := &node
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
		tp.Left = &t
	} else {
		tp.Right = &t
	}
	return &node
}

func copy(node TreeNode) *TreeNode {
	t := TreeNode{node.Val, nil, nil}
	t.Val = node.Val
	if node.Left != nil {
		t.Left = copy(*node.Left)

	}
	if node.Right != nil {
		t.Right = copy(*node.Right)
	}
	return &t
}

func errogenerateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{1, nil, nil}}
	}
	rres := errogenerateTrees(n - 1)
	res := []*TreeNode{}
	lr := len(rres)
	for i := 0; i < lr; i++ {
		res = append(res, add(*copy(*rres[i]), TreeNode{n, nil, nil}))
		res = append(res, add(TreeNode{n, nil, nil}, *copy(*rres[i])))
	}
	return res
}
func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{1, nil, nil}}
	}
	rres := generateTrees(n - 1)
	res := []*TreeNode{}
	lr := len(rres)
	for i := 0; i < lr; i++ {
		res = append(res, add(*copy(*rres[i]), TreeNode{n, nil, nil}))
		res = append(res, add(TreeNode{n, nil, nil}, *copy(*rres[i])))
	}
	return res
}

func main() {

}

/*
1
2
5
14
*/
