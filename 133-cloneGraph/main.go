package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	mp := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		var new *Node
		if mp[node] == nil {
			new = &Node{node.Val, []*Node{}}
			mp[node] = new
		} else {
			return mp[node]
		}
		for i := 0; i < len(node.Neighbors); i++ {
			new.Neighbors = append(new.Neighbors, dfs(node.Neighbors[i]))
		}
		return new
	}
	return dfs(node)
}

func main() {

}
