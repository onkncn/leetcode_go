package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numsNoEmpty(lists []*ListNode) int {
	nums := 0
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			nums++
			if nums == 2 {
				return 2
			}
		}
	}
	return nums
}
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	t := new(ListNode)
	res := t
	for numsNoEmpty(lists) == 2 {
		index := -1
		min := 11111
		for i := 0; i < l; i++ {
			if lists[i] != nil && lists[i].Val < min {
				min = lists[i].Val
				index = i
			}
		}
		t.Next = lists[index]
		lists[index] = lists[index].Next
		t = t.Next
	}
	if numsNoEmpty(lists) == 1 {
		for i := 0; i < l; i++ {
			if lists[i] != nil {
				t.Next = lists[i]
				lists[i] = lists[i].Next
				t = t.Next
				break
			}
		}
	}
	return res.Next
}

func main() {

}
