package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createrTree(a []int) *TreeNode {
	l := len(a)
	if l == 0 {
		return nil
	}
	var help func(a []int, index int) *TreeNode
	help = func(a []int, index int) *TreeNode {
		if a[index] == -1 {
			return nil
		}
		t := &TreeNode{a[index], nil, nil}
		if index*2+1 < len(a) {
			t.Left = help(a, index*2+1)
		}
		if index*2+2 < len(a) {
			t.Right = help(a, index*2+2)
		}
		return t
	}
	return help(a, 0)
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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
func reverse(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	newhead := reverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newhead
}
func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 5)
	wg.Add(5)
	rand.Seed(time.Now().UnixNano())
	var create = func() {
		ch <- rand.Intn(100)
	}
	var printer = func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}
	for i := 0; i < 5; i++ {
		go create()
		go printer()
	}
	wg.Wait()
}
