package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return [][]int{}
	}
	res = [][]int{}
	queue := []*TreeNode{root}
	lq := 1
	for lq != 0 {
		defer func(a []*TreeNode) {
			t := []int{}
			for _, c := range a {
				t = append(t, c.Val)
			}
			res = append(res, t)
		}(queue)
		for i := 0; i < lq; i++ {
			t := queue[i]
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
		}
		queue = queue[lq:]
		lq = len(queue)
	}
	return res
}

func main() {

}
