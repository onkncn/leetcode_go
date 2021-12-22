package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := new(ListNode)
	head := t
	c := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		t.Next = &ListNode{(v1 + v2 + c) % 10, nil}
		t = t.Next
		c = (v1 + v2 + c) / 10
	}
	if c != 0 {
		t.Next = &ListNode{c, nil}
	}
	return head.Next
}

func main() {

}
