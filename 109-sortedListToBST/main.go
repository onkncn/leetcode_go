package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}
	f, s := &ListNode{0, head}, head
	for s != nil && s.Next != nil {
		s = s.Next.Next
		f = f.Next
	}
	head2 := f.Next.Next
	t := f.Next.Val
	f.Next = nil
	return &TreeNode{t, sortedListToBST(head), sortedListToBST(head2)}
} //1 2 3 4

func main() {

}
