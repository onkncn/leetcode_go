package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
	h := helpSort(head)
	return h
}
func merge(l1, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
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
	return head.Next
}
func helpSort(head *ListNode) *ListNode {
	s, q := head, head.Next
	if q == nil {
		return head
	}
	if q.Next != nil && q.Next.Next != nil {
		q = q.Next.Next
		s = s.Next
	}
	l2 := s.Next
	s.Next = nil
	h1 := helpSort(head)
	h2 := helpSort(l2)
	return merge(h1, h2)
}

func main() {
	l := createrList([]int{5, 1, 2, 3, 4})
	printList(sortList(l))
}
