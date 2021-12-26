package main

func numFriendRequests(ages []int) int {
	age := [121]int{}
	for _, a := range ages {
		age[a]++
	}
	sum := 0
	for x := 15; x <= 120; x++ {
		t := age[x]
		if t == 0 {
			continue
		}
		y := x
		if y <= 100 || x >= 100 {
			if x == y {
				sum += age[x] * (age[x] - 1)
			}
			y--
			for 2*y > x+14 {
				sum += age[x] * age[y]
				y--
			}
		}
	}
	return sum
}

func main() {

}
