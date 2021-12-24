package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	lq := 1
	flag := true
	for lq != 0 {
		t := queue[0].Val
		if flag {
			t--
		} else {
			t++
		}
		for i := 0; i < lq; i++ {
			if (flag && (queue[i].Val%2 == 0 || t >= queue[i].Val)) || (!flag && (queue[i].Val%2 == 1 || t <= queue[i].Val)) {
				return false
			}
			t = queue[i].Val
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
	return true
}

func main() {

}
