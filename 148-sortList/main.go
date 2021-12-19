package main

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

func createrTree(a []int) *TreeNode {
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

func createrList(nums []int) *ListNode {
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

func printList(node *ListNode) {
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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s, q := head, head.Next
	for q.Next != nil && q.Next.Next != nil {
		q = q.Next.Next
		s = s.Next
	}
	l2 := s.Next
	s.Next = nil
	return merge(sortList(head), sortList(l2))
}
func merge(l1, l2 *ListNode) *ListNode {
	res := l1
	if l1.Val > l2.Val {
		res = l2
		l2 = l2.Next
	} else {
		l1 = l1.Next
	}
	head := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Next = l1
			l1 = l1.Next
			res = res.Next
		} else {
			res.Next = l2
			l2 = l2.Next
			res = res.Next
		}
	}
	if l1 != nil {
		res.Next = l1
	}
	if l2 != nil {
		res.Next = l2
	}
	return head
}
func main() {
	l := createrList([]int{5, 1, 2, 3, 4})
	printList(sortList(l))
}
