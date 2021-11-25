package main

import (
	"container/list"
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	ss := strings.Split(path, "/")
	res := ""
	flag := list.List{}
	for i := 0; i < len(ss); i++ {
		if ss[i] != "" {
			if ss[i] == ".." {
				if flag.Back() != nil {
					flag.Remove(flag.Back())
				} else {
					continue
				}
			} else if ss[i] == "." {
				continue
			} else {
				flag.PushBack(ss[i])
			}
		}
	}
	if flag.Front() == nil {
		return "/"
	}
	for i := flag.Front(); i != nil; i = i.Next() {
		res += "/" + i.Value.(string)
	}
	return res
}

func main() {
	fmt.Println(simplifyPath("/home//foo/"))
}
