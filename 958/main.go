package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	l := 1
	last := false
	for l != 0 {
		flag := false
		for i := 0; i < l; i++ {
			t := queue[i]
			if t.Left != nil {
				if flag || last {
					return false
				}
				queue = append(queue, t.Left)
			} else {
				flag = true
			}
			if t.Right != nil {
				if flag || last {
					return false
				}
				queue = append(queue, t.Right)
			} else {
				flag = true
			}
		}
		last = flag
		queue = queue[l:]
		l = len(queue)
	}
	return true
}

/*
 1
 2 3
4 5 6 7
8 9 10 11 12 13 14 15*/
