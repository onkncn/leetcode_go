package main

import (
	"fmt"
	"strconv"
)

func acc(s string) bool {
	ls := len(s)
	if ls == 1 {
		return true
	} else if s[0] != '0' {
		i, _ := strconv.Atoi(s)
		if int(i) <= 255 {
			return true
		}
	}
	return false
}
func acc4(s []string) bool {
	for i := 0; i < 4; i++ {
		if !acc(s[i]) {
			return false
		}
	}
	return true
}
func restoreIpAddresses(s string) []string {
	ls := len(s)
	if ls < 4 || ls > 12 {
		return []string{}
	}
	if len(s) == 4 {
		return []string{string(s[0]) + "." + string(s[1]) + "." + string(s[2]) + "." + string(s[3])}
	}
	res := []string{}
	for i1 := 1; i1 < 4; i1++ {
		for i2 := i1 + 1; i2 < i1+4; i2++ {
			for i3 := i2 + 1; i3 < i2+4 && i3 < ls; i3++ {
				ss := []string{s[0:i1], s[i1:i2], s[i2:i3], s[i3:ls]}
				if acc4(ss) {
					res = append(res, ss[0]+"."+ss[1]+"."+ss[2]+"."+ss[3])
					fmt.Println(res)
				}
			}
		}
	}
	return res
}

//12345678
func main() {
	fmt.Println(restoreIpAddresses("122312"))
}
