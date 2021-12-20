package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createrList(nums []int) *ListNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	t := head
	for i := 1; i < l; i++ {
		t.Next = &ListNode{nums[i], nil}
		t = t.Next
	}
	return head
}

func printList(node *ListNode) {
	start := node
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
		if node == start {
			break
		}
	}
	fmt.Println()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	var reverse func(head *ListNode, times int) *ListNode
	reverse = func(head *ListNode, times int) *ListNode {
		if head == nil || head.Next == nil || times == 1 {
			return head
		}
		newhead := reverse(head.Next, times-1)
		head.Next.Next = head
		head.Next = nil
		return newhead
	}
	l := 0
	t, start, ans, laststart := head, head, head, head
	for t != nil {
		next := t.Next
		if l%k == k-1 {
			if l/k == 0 {
				ans = reverse(start, k)
				start.Next = next
			} else {
				laststart.Next = reverse(start, k)
				start.Next = next
			}
			laststart = start
			start = next
			//printList(ans)
		}
		l++
		t = next
	}
	return ans
}
func main() {
	nums := []int{1, 2}
	printList(reverseKGroup(createrList(nums), 2))
}
