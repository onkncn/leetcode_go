package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	res := []*ListNode{}
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	l := len(res)
	for i := 0; i < l/2; i++ {
		res[i].Next = res[l-1-i]
		res[l-1-i].Next = res[i+1]
	}
	res[l/2].Next = nil
}

func main() {

}
