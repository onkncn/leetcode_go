package main

import "sort"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func mymax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func mymin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func findRadius(houses []int, heaters []int) int {
	lhouse := len(houses)
	lheater := len(heaters)
	sort.Ints(heaters)
	sort.Ints(houses)
	heaters = append(heaters, heaters[lheater-1])
	h := 1
	max := abs(0)
	for i := 0; i < lhouse; i++ {
		for h < lheater && heaters[h] < houses[i] {
			h++
		}
		max = mymax(mymin(abs(houses[i]-heaters[h]), abs(houses[i]-heaters[h-1])), max)
	}
	return max
}

func main() {

}
