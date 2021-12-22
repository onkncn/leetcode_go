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
	queue := []*Node{root}
	lq := 1
	for lq != 0 {
		for i := 0; i < lq; i++ {
			if queue[i] != nil {
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
				if i != lq-1 {
					queue[i].Next = queue[i+1]
				}
			}
		}
		queue = queue[lq:]
		lq = len(queue)
	}
	return root
}

func main() {

}
