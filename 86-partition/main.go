package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	big := ListNode{0, nil}
	sma := ListNode{0, nil}
	tb := &big
	ts := &sma
	for head != nil {
		if head.Val < x {
			ts.Next = head
			ts = ts.Next
		} else {
			tb.Next = head
			tb = tb.Next
		}
		head = head.Next
	}
	tb.Next = nil
	ts.Next = big.Next
	if sma.Next != nil {
		return sma.Next
	}
	return big.Next
}

func main() {
	bg1 := &ListNode{1, nil}
	bg2 := &ListNode{4, nil}
	bg3 := &ListNode{3, nil}
	bg4 := &ListNode{2, nil}
	bg5 := &ListNode{5, nil}
	bg6 := &ListNode{2, nil}
	bg1.Next = bg2
	bg2.Next = bg3
	bg3.Next = bg4
	bg4.Next = bg5
	bg5.Next = bg6
	t := partition(bg1, 3)
	for t != nil {
		fmt.Println(t.Val)
		t = t.Next
	}
}
