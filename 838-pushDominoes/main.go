package main

import "fmt"

func pushDominoes(dominoes string) string {
	l := len(dominoes)
	res := make([]byte, l)
	copy(res, dominoes)
	mpL := []int{}
	mpR := []int{}
	for i := 0; i < l; i++ {
		if dominoes[i] == 'L' {
			mpL = append(mpL, i)
		} else if dominoes[i] == 'R' {
			mpR = append(mpR, i)
		}
	}
	ll, lr := len(mpL), len(mpR)
	for ll+lr != 0 {
		view := make([]int, l)
		for i := 0; i < ll; i++ {
			if mpL[i]-1 >= 0 && res[mpL[i]-1] == '.' {
				res[mpL[i]-1] = 'L'
				mpL = append(mpL, mpL[i]-1)
				view[mpL[i]-1] = 1
			}
		}
		for i := 0; i < lr; i++ {
			if mpR[i]+1 < l {
				if res[mpR[i]+1] == '.' {
					res[mpR[i]+1] = 'R'
					mpR = append(mpR, mpR[i]+1)
				} else if view[mpR[i]+1] == 1 {
					res[mpR[i]+1] = '.'
				}
			}
		}
		mpL, mpR = mpL[ll:], mpR[lr:]
		ll, lr = len(mpL), len(mpR)
	}
	return string(res)
}

func main() {
	fmt.Println(pushDominoes("RR.L"))
}
