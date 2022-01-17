package help

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreaterTree(a []int) *TreeNode {
	l := len(a)
	if l == 0 {
		return nil
	}
	var help func(a []int, index int) *TreeNode
	help = func(a []int, index int) *TreeNode {
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
	return help(a, 0)
}

func CreaterList(nums []int) *ListNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	t := head
	for i := 1; i < l; i++ {
		t.Next = &ListNode{nums[i], nil}
		t = t.Next
	}
	return head
}

func PrintList(node *ListNode) {
	start := node
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
		if node == start {
			break
		}
	}
	fmt.Println()
}
