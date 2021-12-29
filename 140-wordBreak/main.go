package main

import "fmt"

type Tree struct {
	End  bool
	Next []*Tree
}

func insert(head *Tree, s string) {
	l := len(s)
	for i := 0; i < l; i++ {
		if head.Next[s[i]-'a'] == nil {
			head.Next[s[i]-'a'] = &Tree{false, make([]*Tree, 26)}
		}
		head = head.Next[s[i]-'a']
	}
	head.End = true
}
func find(head *Tree, s string) []string {
	res := []string{}
	l := len(s)
	root := head
	i := 0
	for i = 0; i < l; i++ {
		if root.End {
			t := find(head, s[i:])
			for _, temps := range t {
				res = append(res, s[:i]+" "+temps)
			}
		}
		if root.Next[s[i]-'a'] == nil {
			break
		}
		root = root.Next[s[i]-'a']
	}
	if i == l && root.End {
		res = append(res, s)
	}
	return res
}
func wordBreak(s string, wordDict []string) []string {
	head := &Tree{false, make([]*Tree, 26)}
	for _, s := range wordDict {
		insert(head, s)
	}
	return find(head, s)
}

func main() {
	res := wordBreak("aabb", []string{"a", "b", "aa", "ab", "bb"})
	for _, s := range res {
		fmt.Println(s)
	}
}
