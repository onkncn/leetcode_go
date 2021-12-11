package main

import (
	"fmt"
)

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	for i := 0; i < len(words); {
		tl := -1
		ts := []string{}
		j := i
		for ; j < len(words); j++ {
			tl += len(words[j]) + 1
			if tl <= maxWidth {
				ts = append(ts, words[j])
			} else {
				break
			}
		}
		i = j
		lts := len(ts) //单词数
		tl -= lts      //减去计算时每个单词后面的空格
		s := ""
		if j < len(words) {
			tl = tl - len(words[j]) //去掉最后那个单词
		} else {
			s = ts[0]
			for j := 1; j < len(ts); j++ {
				s += " " + ts[j]
			}
			for j := len(s); j < maxWidth; j++ {
				s += " "
			}
			res = append(res, s)
			return res
		}
		ns := maxWidth - tl //空格总数
		fmt.Println(ts)
		if lts == 1 {
			s += ts[0]
			for j := len(s); j < maxWidth; j++ {
				s += " "
			}
			res = append(res, s)
			continue
		}
		times := ns / (lts - 1)
		for j := 0; j < 2*lts-1; j++ {
			if j%2 == 0 {
				s += ts[j/2]
			} else {
				for k := 0; k < times; k++ {
					s += " "
				}
				if ns%(lts-1) != 0 {
					s += " "
					ns--
				}
			}
		}
		res = append(res, s)
	}
	return res
}

func main() {
	ss := fullJustify([]string{"what", "you", "ca", "do", "for", "your", "country"}, 16)
	for _, s := range ss {
		fmt.Println(s)
	}
}
