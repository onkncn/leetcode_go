package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for p != nil && p.Next != nil {
		for p.Next != nil && p.Val == p.Next.Val {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return head
}
func main() {

}
