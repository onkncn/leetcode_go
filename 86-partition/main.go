package main

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
	ts.Next = big.Next
	if sma.Next != nil {
		return sma.Next
	}
	return big.Next
}

func main() {

}
