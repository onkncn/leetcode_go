package main

import (
	"fmt"
	"math/rand"
	//"github.com/Leetcode/help"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() int {
	index := 1
	res, temp := this.head, this.head
	for temp.Next != nil {
		if rand.Intn(index+1) == index {
			res = temp.Next
		}
		temp = temp.Next
	}
	return res.Val
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(2))
	}
}
