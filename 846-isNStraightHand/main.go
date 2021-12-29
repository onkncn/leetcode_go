package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	mp := make(map[int]int)
	l := len(hand)
	if l%groupSize != 0 {
		return false
	}
	for i := 0; i < l; i++ {
		mp[hand[i]]++
	}
	ks := []int{}
	for k, _ := range mp {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	//fmt.Println(ks)
	for _, i := range ks {
		for mp[i] > 0 {
			for j := i; j < i+groupSize; j++ {
				mp[j]--
				if mp[j] < 0 {
					return false
				}
			}
		}
	}
	return true
}

func main() {

}
