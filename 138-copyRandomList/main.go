package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	var dfs func(head *Node) *Node
	dfs = func(head *Node) *Node {
		if head == nil {
			return nil
		}
		if mp[head] != nil {
			return mp[head]
		}
		new := &Node{head.Val, nil, nil}
		mp[head] = new
		new.Next = dfs(head.Next)
		new.Random = dfs(head.Random)
		return new
	}
	return dfs(head)
}
func main() {
	node := &Node{0, &Node{0, nil, nil}, &Node{0, nil, nil}}
	node.Next.Next = node
	copyRandomList(node)
}
