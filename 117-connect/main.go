package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0, 4096)
	queue = append(queue, root)
	lq := 1
	for lq != 0 {
		flag := false
		t := root
		for i := 0; i < lq; i++ {
			if queue[i] != nil {
				if flag {
					t.Next = queue[i]
					t = queue[i]
				} else {
					flag = true
					t = queue[i]
				}
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[lq:]
		lq = len(queue)
	}
	return root
}

func main() {

}
