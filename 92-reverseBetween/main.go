package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	if head.Next == tail {
		tail.Next = head
		head.Next = nil
		return tail, head
	}

	th, tt := reverse(head.Next, tail)
	tt.Next = head
	head.Next = nil
	return th, head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	myhead := ListNode{0, head}
	p := &myhead
	for i := 0; i < left-1; i++ {
		p = p.Next
	}
	q := p
	for i := left; i < right+1; i++ {
		q = q.Next
	}
	rtn := q.Next
	h, t := reverse(p.Next, q)
	p.Next = h
	t.Next = rtn
	return myhead.Next
}

// x 1,2,3,4,5

func main() {
	bg1 := &ListNode{1, nil}
	bg2 := &ListNode{2, nil}
	bg3 := &ListNode{3, nil}
	bg4 := &ListNode{4, nil}
	bg5 := &ListNode{5, nil}
	bg1.Next = bg2
	bg2.Next = bg3
	bg3.Next = bg4
	bg4.Next = bg5
	t := reverseBetween(bg1, 2, 4)
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}
}
