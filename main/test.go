package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createrTree(a []int) *TreeNode {
	l := len(a)
	if l == 0 {
		return nil
	}
	return help(a, 0)
}
func help(a []int, index int) *TreeNode {
	if a[index] == -1 {
		return nil
	}
	t := &TreeNode{a[index], nil, nil}
	if index*2+1 < len(a) {
		t.Left = help(a, index*2+1)
	}
	if index*2+2 < len(a) {
		t.Right = help(a, index*2+2)
	}
	return t
}

func main() {
	fmt.Println(math.Atan2(-0.01, -1) / math.Pi * 180)
}
