package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func bfs(h *Node, i int) int {
	maxi := i
	for _, n := range h.Children {
		maxi = max(maxi, bfs(n, i+1))
	}
	return maxi
}
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return bfs(root, 1)
}
func main() {
	fmt.Println("123")
}
