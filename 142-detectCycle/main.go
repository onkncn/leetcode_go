package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	mp := make(map[*ListNode]bool)
	t := head
	for t != nil && !mp[t] {
		mp[t] = true
		t = t.Next
	}
	return t
}

func main() {

}
