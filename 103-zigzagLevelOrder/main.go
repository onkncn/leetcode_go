package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := []*TreeNode{root}
	lq := 1
	flag := true
	for lq != 0 {
		t := []int{}
		if flag {
			for i := 0; i < lq; i++ {
				t = append(t, queue[i].Val)
			}
		} else {
			for i := lq - 1; i >= 0; i-- {
				t = append(t, queue[i].Val)
			}
		}
		res = append(res, t)
		for i := 0; i < lq; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		flag = !flag
		queue = queue[lq:]
		lq = len(queue)
	}
	return res
}

func main() {

}
