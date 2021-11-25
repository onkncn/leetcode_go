package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	llen := 1
	t := head
	if head == nil || head.Next == nil {
		return head
	}
	for t.Next != nil {
		llen++
		t = t.Next
	}
	k = llen - k%llen
	if k == 0 {
		return head
	}
	p := head
	for i := 0; i < k-1; i++ {
		p = p.Next
	}
	q := p.Next
	for q.Next != nil {
		q = q.Next
	}
	q.Next = head
	res := p.Next
	p.Next = nil
	return res
}
func main() {
	fmt.Println("")
}
