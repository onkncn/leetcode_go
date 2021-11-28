package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	myhead := &ListNode{-111, head}
	p := myhead
	for p.Next != nil && p.Next.Next != nil {
		t := p.Next.Val
		t2 := p.Next.Next.Val
		if t == t2 {
			tp := p.Next
			for tp != nil && tp.Val == t {
				tp = tp.Next
			}
			if tp == nil {
				p.Next = nil
				return myhead.Next
			}
			p.Next = tp
		} else {
			p = p.Next
		}
	}
	return myhead.Next
}

func main() {

}
